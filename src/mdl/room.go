package mdl

import (
	"encoding/json"
	"fmt"

	"github.com/matrix-org/gomatrix"
	"github.com/neilalexander/mdl/src/mdl/base"
	lua "github.com/yuin/gopher-lua"
)

const EVENT_TYPE = "com.github.neilalexander.mdl.source"

type Room struct {
	client *Client
	roomID string
	state  *Lua
}

func (r *Room) NewLua() (*Lua, error) {
	vm := &Lua{
		room: r,
		state: lua.NewState(lua.Options{
			MinimizeStackMemory: true,
			SkipOpenLibs:        true,
		}),
		modules: defaultModules,
	}

	lua.OpenString(vm.state)
	lua.OpenMath(vm.state)
	lua.OpenTable(vm.state)

	vm.state.SetGlobal("require", vm.state.NewFunction(vm.require))
	vm.state.SetGlobal("load", vm.state.NewFunction(vm.load))

	vm.state.SetGlobal("print", vm.state.NewFunction(base.Print))

	return vm, nil
}

func (r *Room) Event(eventID string) (*Source, error) {
	var event gomatrix.Event
	err := r.client.client.Event(r.roomID, eventID, &event)
	switch e := err.(type) {
	case nil:
		if event.Type != EVENT_TYPE {
			return nil, fmt.Errorf("unexpected EVENT_TYPE: %q", event.Type)
		}
	case gomatrix.HTTPError:
		return nil, fmt.Errorf("r.client.client.Event HTTP code %d: %s\nContents: %s", e.Code, e.Message, e.Contents)
	default:
		return nil, fmt.Errorf("r.client.client.Event: %w", err)
	}
	return NewSourceFromEvent(event)
}

func (r *Room) StateEvent(stateKey string) (*Source, error) {
	var event map[string]interface{}
	err := r.client.client.StateEvent(r.roomID, EVENT_TYPE, stateKey, &event)
	switch e := err.(type) {
	case nil:
		j, err := json.Marshal(event)
		if err != nil {
			return nil, fmt.Errorf("json.Marshal: %w", err)
		}
		return NewSourceFromJSON(j)
	case gomatrix.HTTPError:
		return nil, fmt.Errorf("r.client.client.StateEvent HTTP code %d: %s\nContents: %s", e.Code, e.Message, e.Contents)
	default:
		return nil, fmt.Errorf("r.client.client.StateEvent: %w", err)
	}
}
