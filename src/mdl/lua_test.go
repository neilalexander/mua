package mdl

/*
func TestInterpreter(t *testing.T) {
	script := `
		foo = require("test")
		foo.print("This module is called '" .. foo.name .. "' and it is from '" .. foo.source .. "'")
	`

	var r *Room
	lua, err := r.NewLua()
	if err != nil {
		t.Fatalf("Unable to start Lua: %s", err)
	}

	if err = lua.state.DoString(script); err != nil {
		t.Fatalf("Failed to run test script: %s", err)
	}
}
*/
