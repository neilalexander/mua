package mua

import (
	"errors"

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
		defaultModules,
	}

	lua.StringOpen(L.State)
	lua.MathOpen(L.State)
	lua.TableOpen(L.State)
	lua.Bit32Open(L.State)

	L.registerEventType()
	L.registerBaseFunctions()

	return L, nil
}

func (vm *Lua) Execute(source string) error {
	return lua.DoString(vm.State, source)
}

func (vm *Lua) ExecuteFile(file string) error {
	return lua.DoFile(vm.State, file)
}

func (vm *Lua) requestModule(moduleName string) (*LuaModule, error) {
	return nil, errors.New("requestModule not implemented")
}
