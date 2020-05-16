package mdl

import (
	"github.com/neilalexander/mdl/src/mdl/base"
	lua "github.com/yuin/gopher-lua"
)

type LuaModule struct {
	table map[string]lua.LGFunction
}

var defaultModules = map[string]LuaModule{
	"test": {
		table: map[string]lua.LGFunction{
			"print": base.Print,
		},
	},
}
