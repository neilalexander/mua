package mdl

import (
	"errors"
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

type Lua struct {
	room    *Room
	state   *lua.LState
	modules map[string]LuaModule
}

type LuaModule struct {
	table map[string]lua.LGFunction
}

var defaultModules = map[string]LuaModule{
	"test": {
		table: map[string]lua.LGFunction{
			//"print": print,
		},
	},
}

func (vm *Lua) Execute(source string) error {
	return vm.state.DoString(source)
}

func (vm *Lua) require(L *lua.LState) int {
	moduleName := L.ToString(1)
	module, ok := vm.modules[moduleName]
	if ok {
		table := L.SetFuncs(L.NewTable(), module.table)
		L.SetField(table, "name", lua.LString(moduleName))
		L.Push(table)
		return 1
	}
	L.ArgError(1, fmt.Sprintf("module %q does not exist", moduleName))
	return 0
}

func (vm *Lua) requestModule(moduleName string) (*LuaModule, error) {
	return nil, errors.New("requestModule not implemented")
}

func (vm *Lua) load(L *lua.LState) int {
	moduleName := L.ToString(1)

	var src *Source
	var err error

	if len(moduleName) == 0 {
		L.ArgError(1, "must specify an event ID or state key")
		return 0
	}

	switch moduleName[0] {
	case '$':
		src, err = vm.room.Event(moduleName)
	default:
		src, err = vm.room.StateEvent(moduleName)
	}
	if err != nil {
		L.RaiseError("attempt to load %q failed: %s", moduleName, err)
		return 0
	}

	if err != nil {
		L.RaiseError("vm.loadModule: %s", err)
		return 0
	}

	if err := L.DoString(string(src.Source)); err != nil {
		L.RaiseError("L.DoString: %s", err)
	}

	return 0
}

func (vm *Lua) print(L *lua.LState) int {
	str := L.CheckString(1)
	fmt.Printf("\033[1;32m%s\033[0m\n", str)
	return 0
}
