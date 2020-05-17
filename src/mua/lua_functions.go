package mua

import (
	"fmt"

	"github.com/Shopify/go-lua"
)

func (vm *Lua) registerBaseFunctions() {
	vm.PushGoFunction(vm.print)
	vm.SetGlobal("print")

	vm.PushGoFunction(vm.require)
	vm.SetGlobal("require")

	vm.PushGoFunction(vm.importevent)
	vm.SetGlobal("importevent")

	vm.PushGoFunction(vm.importstate)
	vm.SetGlobal("importstate")
}

func (vm *Lua) require(L *lua.State) int {
	moduleName, ok := L.ToString(1)
	if !ok {
		L.PushString("Expecting (module)")
		L.Error()
		return 0
	}

	module, ok := vm.modules[moduleName]
	if !ok {
		L.PushString(fmt.Sprintf("Module %q does not exist", moduleName))
		L.Error()
		return 0
	}

	lua.NewLibrary(vm.State, module.RegistryFunctions())
	vm.State.SetGlobal(moduleName)

	return 1
}

func (vm *Lua) importevent(L *lua.State) int {
	roomID, ok := L.ToString(1)
	if !ok {
		L.PushString("Expecting roomid as string in first parameter")
		L.Error()
		return 0
	}

	eventID, ok := L.ToString(2)
	if !ok {
		L.PushString("Expecting eventid as string in second parameter")
		L.Error()
		return 0
	}

	room, ok := vm.client.rooms[roomID]
	if !ok {
		L.PushString("Room out of scope")
		L.Error()
		return 0
	}

	src, err := room.SourceFromEvent(eventID)
	if err != nil {
		L.PushString(fmt.Sprintf("Attempt to load %q failed: %s", eventID, err))
		L.Error()
		return 0
	}

	if err := vm.Execute(string(src.Source)); err != nil {
		L.PushString(fmt.Sprintf("L.DoString: %s", err))
		L.Error()
		return 0
	}

	return 0
}

func (vm *Lua) importstate(L *lua.State) int {
	roomID, ok := L.ToString(1)
	if !ok {
		L.PushString("Expecting roomid as string in first parameter")
		L.Error()
		return 0
	}

	eventType, ok := L.ToString(2)
	if !ok {
		L.PushString("Expecting eventtype as string in second parameter")
		L.Error()
		return 0
	}

	stateKey, ok := L.ToString(3)
	if !ok {
		L.PushString("Expecting statekey as string in third parameter")
		L.Error()
		return 0
	}

	room, ok := vm.client.rooms[roomID]
	if !ok {
		L.PushString("Room out of scope")
		L.Error()
		return 0
	}

	src, err := room.SourceFromStateEvent(eventType, stateKey)
	if err != nil {
		L.PushString(fmt.Sprintf("Attempt to load %q %q failed: %s", eventType, stateKey, err))
		L.Error()
		return 0
	}

	if err := vm.Execute(string(src.Source)); err != nil {
		L.PushString(fmt.Sprintf("L.DoString: %s", err))
		L.Error()
		return 0
	}

	return 0
}

func (vm *Lua) print(L *lua.State) int {
	str, ok := L.ToString(1)
	if !ok {
		L.PushString(fmt.Sprintf("Expected (string)"))
		L.Error()
		return 0
	}
	fmt.Printf("\033[1;32m%s\033[0m\n", str)
	return 0
}
