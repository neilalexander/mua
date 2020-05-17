package mua

import (
	"encoding/json"
	"fmt"
	"reflect"

	lua "github.com/Shopify/go-lua"
	"github.com/matrix-org/gomatrix"
)

const EVENT_TYPE = "mua.source"

func (vm *Lua) registerEventType() {
	if !lua.NewMetaTable(vm.State, "event") {
		return
	}

	vm.PushString("__index")
	vm.PushValue(-2)
	vm.SetTable(-3)

	for _, f := range []lua.RegistryFunction{
		{"new", vm.newEvent},
		{"state_key", eventGetSetStateKey},
		{"sender", eventGetSetSender},
		{"type", eventGetSetType},
		{"origin_server_ts", eventGetSetOriginServerTS},
		{"event_id", eventGetSetEventID},
		{"room_id", eventGetSetRoomID},
		{"redacts", eventGetSetRedacts},
		//{"unsigned", eventGetSetUnsigned},
		//{"content", eventGetSetContent},
		//{"prev_content", eventGetSetPrevContent},
		{"json", eventGetJSON},
	} {
		vm.PushGoFunction(f.Function)
		vm.SetField(-2, f.Name)
	}

	vm.SetGlobal("event")
}

func (vm *Lua) newEvent(L *lua.State) int {
	var err error
	var event *gomatrix.Event

	if L.Top() != 2 {
		L.PushString("expected (roomid, eventid)")
		L.Error()
		return 0
	}

	roomID, ok1 := L.ToString(1)
	eventID, ok2 := L.ToString(2)
	if !ok1 || !ok2 {
		L.PushString("expected (roomid, eventid)")
		L.Error()
		return 0
	}

	room, ok3 := vm.client.rooms[roomID]
	if !ok3 {
		L.PushString("room is not in scope")
		L.Error()
		return 0
	}

	if event, err = room.Event(eventID); err != nil {
		L.PushString(fmt.Sprintf("Failed to get event %q: %s", eventID, err))
		L.Error()
		return 0
	}

	L.PushUserData(event)
	lua.SetMetaTableNamed(L, "event")
	return 1
}

func checkEvent(L *lua.State) *gomatrix.Event {
	ok := L.IsUserData(1)
	if !ok {
		L.PushString("Expected userdata")
		L.Error()
		return nil
	}

	if v, ok := L.ToUserData(1).(*gomatrix.Event); ok {
		return v
	}

	L.PushString("Expected event")
	L.Error()
	return nil
}

func eventGetSetReflected(L *lua.State, key string) int {
	p := checkEvent(L)
	ps := reflect.ValueOf(p)
	field := ps.Elem().FieldByName(key)

	switch L.Top() {
	case 2:
		switch true {
		case L.IsString(1):
			value, _ := L.ToString(1)
			field.SetString(value)
		case L.IsBoolean(1):
			value := L.ToBoolean(1)
			field.SetBool(value)
		case L.IsNumber(1):
			value, _ := L.ToNumber(1)
			field.SetUint(uint64(value))
		case L.IsString(1):
			value, _ := L.ToString(1)
			field.SetString(value)
		default:
			L.PushString("Unsupported type")
			L.Error()
			return 0
		}
		fallthrough
	case 1:
		L.PushString(field.String())
		return 1
	default:
		L.PushString("Expected 1 or 2")
		L.Error()
		return 0
	}
}

func eventGetSetStateKey(L *lua.State) int {
	p := checkEvent(L)
	switch L.Top() {
	case 2:
		if stateKey, ok := L.ToString(2); ok {
			p.StateKey = &stateKey
		}
		fallthrough
	case 1:
		if p.StateKey != nil {
			L.PushString(*p.StateKey)
		} else {
			L.PushString("")
		}
		return 1
	default:
		L.PushString("Expected 1 or 2")
		L.Error()
		return 0
	}
}

func eventGetSetSender(L *lua.State) int {
	return eventGetSetReflected(L, "Sender")
}

func eventGetSetType(L *lua.State) int {
	return eventGetSetReflected(L, "Type")
}

func eventGetSetOriginServerTS(L *lua.State) int {
	return eventGetSetReflected(L, "Timestamp")
}

func eventGetSetEventID(L *lua.State) int {
	return eventGetSetReflected(L, "ID")
}

func eventGetSetRoomID(L *lua.State) int {
	return eventGetSetReflected(L, "RoomID")
}

func eventGetSetRedacts(L *lua.State) int {
	return eventGetSetReflected(L, "Redacts")
}

func eventGetJSON(L *lua.State) int {
	p := checkEvent(L)
	j, err := json.Marshal(p)
	if err != nil {
		L.PushString(fmt.Sprintf("Failed to marshal: %s", err))
		L.Error()
		return 0
	}
	L.PushString(string(j))
	return 1
}
