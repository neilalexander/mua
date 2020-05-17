package mua

import (
	"os"

	"github.com/Shopify/go-lua"
)

var moduleMua = map[string]lua.Function{
	"test": func(L *lua.State) int {
		L.Global("print")
		L.PushString("Test successful")
		L.Call(1, 0)
		return 0
	},
	"dump": func(L *lua.State) int {
		lua.DoString(L, `do print("HELLO") end`)
		L.Dump(os.Stdout)
		return 0
	},
}
