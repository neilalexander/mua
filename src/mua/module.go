package mua

import (
	"github.com/Shopify/go-lua"
)

type LuaModule struct {
	table map[string]lua.Function
}

func (m *LuaModule) RegistryFunctions() []lua.RegistryFunction {
	var registry []lua.RegistryFunction
	for k, v := range m.table {
		registry = append(registry, lua.RegistryFunction{
			Name:     k,
			Function: v,
		})
	}
	return registry
}

var defaultModules = map[string]LuaModule{
	"mua": {
		table: moduleMua,
	},
}
