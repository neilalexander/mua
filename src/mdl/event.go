package mdl

/*
const luaEventTypeName = "event"

func (vm *Lua) registerEventType() {
	mt := vm.state.NewTypeMetatable(luaEventTypeName)
	vm.state.SetGlobal("event", mt)
	vm.state.SetField(mt, "new", vm.state.NewFunction(vm.newEvent))
	vm.state.SetField(mt, "__index", vm.state.SetFuncs(
		vm.state.NewTable(),
		map[string]lua.LGFunction{
			"state_key":        eventGetSetStateKey,
			"sender":           eventGetSetSender,
			"type":             eventGetSetType,
			"origin_server_ts": eventGetSetOriginServerTS,
			"event_id":         eventGetSetEventID,
			"room_id":          eventGetSetRoomID,
			"redacts":          eventGetSetRedacts,
			"unsigned":         eventGetSetUnsigned,
			"content":          eventGetSetContent,
			"prev_content":     eventGetSetPrevContent,
			"json":             eventGetJSON,
		},
	))
}

func (vm *Lua) newEvent(L *lua.LState) int {
	event := &gomatrix.Event{}
	if L.GetTop() == 1 {
		if eventID := L.CheckString(1); eventID != "" {
			if err := vm.room.client.client.Event(vm.room.roomID, eventID, &event); err != nil {
				L.RaiseError("failed to get event %q: %s", eventID, err)
			}
		}
	}
	ud := L.NewUserData()
	ud.Value = event
	L.SetMetatable(ud, L.GetTypeMetatable(luaEventTypeName))
	L.Push(ud)
	return 1
}

func checkEvent(L *lua.LState) *gomatrix.Event {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*gomatrix.Event); ok {
		return v
	}
	L.ArgError(1, "event expected")
	return nil
}

func eventGetSetStateKey(L *lua.LState) int {
	p := checkEvent(L)
	if L.GetTop() == 2 {
		stateKey := L.CheckString(2)
		p.StateKey = &stateKey
		return 0
	}
	if p.StateKey == nil {
		L.Push(lua.LString(""))
		return 1
	}
	L.Push(lua.LString(*p.StateKey))
	return 1
}

func eventGetSetSender(L *lua.LState) int {
	p := checkEvent(L)
	if L.GetTop() == 2 {
		p.Sender = L.CheckString(2)
		return 0
	}
	L.Push(lua.LString(p.Sender))
	return 1
}

func eventGetSetType(L *lua.LState) int {
	p := checkEvent(L)
	if L.GetTop() == 2 {
		p.Type = L.CheckString(2)
		return 0
	}
	L.Push(lua.LString(p.Type))
	return 1
}

func eventGetSetOriginServerTS(L *lua.LState) int {
	p := checkEvent(L)
	if L.GetTop() == 2 {
		p.Timestamp = L.CheckInt64(2)
		return 0
	}
	L.Push(lua.LNumber(p.Timestamp))
	return 1
}

func eventGetSetEventID(L *lua.LState) int {
	p := checkEvent(L)
	if L.GetTop() == 2 {
		p.ID = L.CheckString(2)
		return 0
	}
	L.Push(lua.LString(p.ID))
	return 1
}

func eventGetSetRoomID(L *lua.LState) int {
	p := checkEvent(L)
	if L.GetTop() == 2 {
		p.RoomID = L.CheckString(2)
		return 0
	}
	L.Push(lua.LString(p.RoomID))
	return 1
}

func eventGetSetRedacts(L *lua.LState) int {
	p := checkEvent(L)
	if L.GetTop() == 2 {
		p.Redacts = L.CheckString(2)
		return 0
	}
	L.Push(lua.LString(p.Redacts))
	return 1
}

func eventGetSetUnsigned(L *lua.LState) int {
	p := checkEvent(L)
	if L.GetTop() == 2 {
		in := []byte(L.CheckString(2))
		if err := json.Unmarshal(in, &p.Unsigned); err != nil {
			L.ArgError(2, err.Error())
		}
		return 0
	}
	j, err := json.Marshal(p.Unsigned)
	if err != nil {
		L.RaiseError(err.Error())
	}
	L.Push(lua.LString(j))
	return 1
}

func eventGetSetContent(L *lua.LState) int {
	p := checkEvent(L)
	if L.GetTop() == 2 {
		in := []byte(L.CheckString(2))
		if err := json.Unmarshal(in, &p.Content); err != nil {
			L.ArgError(2, err.Error())
		}
		return 0
	}
	j, err := json.Marshal(p.Content)
	if err != nil {
		L.RaiseError(err.Error())
	}
	L.Push(lua.LString(j))
	return 1
}

func eventGetSetPrevContent(L *lua.LState) int {
	p := checkEvent(L)
	if L.GetTop() == 2 {
		in := []byte(L.CheckString(2))
		if err := json.Unmarshal(in, &p.PrevContent); err != nil {
			L.ArgError(2, err.Error())
		}
		return 0
	}
	j, err := json.Marshal(p.PrevContent)
	if err != nil {
		L.RaiseError(err.Error())
	}
	L.Push(lua.LString(j))
	return 1
}

func eventGetJSON(L *lua.LState) int {
	p := checkEvent(L)
	j, err := json.Marshal(p)
	if err != nil {
		L.RaiseError(err.Error())
	}
	L.Push(lua.LString(j))
	return 1
}
*/
