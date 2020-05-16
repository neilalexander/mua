package base

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

func Print(L *lua.LState) int {
	str := L.ToString(1)
	fmt.Printf("\033[1;32m%s\033[0m\n", str)
	return 0
}
