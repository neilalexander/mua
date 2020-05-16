package mdl

import (
	"errors"
	"fmt"

	"github.com/Shopify/go-lua"
)

type Lua struct {
	*lua.State
	client  *Client
	modules map[string]LuaModule
}

func NewLua(client *Client) (*Lua, error) {
	L := &Lua{
		lua.NewState(),
		client,
		map[string]LuaModule{
			"mua": {
				table: map[string]lua.Function{},
			},
		},
	}

	//lua.BaseOpen(L.State)
	lua.StringOpen(L.State)
	lua.MathOpen(L.State)
	lua.TableOpen(L.State)
	lua.Bit32Open(L.State)

	//vm.registerEventType()
	L.PushGoFunction(L.print)
	L.SetGlobal("print")

	L.PushGoFunction(L.require)
	L.SetGlobal("require")

	L.PushGoFunction(L.importlua)
	L.SetGlobal("import")

	return L, nil
}

func (vm *Lua) Execute(source string) error {
	return lua.DoString(vm.State, source)
}

func (vm *Lua) ExecuteFile(file string) error {
	return lua.DoFile(vm.State, file)
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
	return 1
}

func (vm *Lua) requestModule(moduleName string) (*LuaModule, error) {
	return nil, errors.New("requestModule not implemented")
}

func (vm *Lua) importlua(L *lua.State) int {
	if L.Top() < 2 {
		L.PushString("Expecting (roomid, eventid) or (roomid, eventtype, statekey)")
		L.Error()
		return 0
	}

	roomID, ok := L.ToString(1)
	if !ok {
		L.PushString("Expecting roomid as string")
		L.Error()
		return 0
	}

	moduleName, ok := L.ToString(2)
	if !ok {
		L.PushString("Expecting modulename as string")
		L.Error()
		return 0
	}

	var src *Source
	var err error

	switch moduleName[0] {
	case '$':
		src, err = vm.client.rooms[roomID].Event(moduleName)
	default:
		src, err = vm.client.rooms[roomID].StateEvent(moduleName)
	}
	if err != nil {
		L.PushString(fmt.Sprintf("Attempt to load %q failed: %s", moduleName, err))
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
