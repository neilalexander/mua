package mua

import (
	"encoding/json"
	"fmt"

	"github.com/matrix-org/gomatrix"
)

type Room struct {
	client *Client
	roomID string
}

func (c *Client) NewRoom(roomID string) (*Room, error) {
	/*
		resp, err := c.client.JoinRoom(roomID, "", map[string]string{
			"display_name": "Mua",
		})

		switch e := err.(type) {
		case nil:
			roomID = resp.RoomID
		case gomatrix.HTTPError:
			if e.Code == 404 {
				create, createErr := c.client.CreateRoom(&gomatrix.ReqCreateRoom{
					Name:       "Mua Room",
					Visibility: "private",
				})
				if createErr != nil {
					return nil, fmt.Errorf("c.client.CreateRoom: %w", createErr)
				}
				roomID = create.RoomID
			} else {
				return nil, fmt.Errorf("c.client.JoinRoom: %s (%s)", e.Message, e.Contents)
			}
		default:
			return nil, fmt.Errorf("c.client.JoinRoom: %w", err)
		}
	*/

	room := &Room{
		client: c,
		roomID: roomID,
	}

	c.rooms[roomID] = room
	return room, nil
}

func (r *Room) Event(eventID string) (*gomatrix.Event, error) {
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
	return &event, nil
}

func (r *Room) State(eventType, stateKey string) (map[string]interface{}, error) {
	var state map[string]interface{}
	err := r.client.client.StateEvent(r.roomID, eventType, stateKey, &state)
	switch e := err.(type) {
	case nil:
		return state, nil
	case gomatrix.HTTPError:
		return nil, fmt.Errorf("r.client.client.StateEvent HTTP code %d: %s\nContents: %s", e.Code, e.Message, e.Contents)
	default:
		return nil, fmt.Errorf("r.client.client.StateEvent: %w", err)
	}
}

func (r *Room) SourceFromEvent(eventID string) (*Source, error) {
	if event, err := r.Event(eventID); err == nil {
		return NewSourceFromEvent(*event)
	} else {
		return nil, err
	}
}

func (r *Room) SourceFromStateEvent(eventType, stateKey string) (*Source, error) {
	if state, err := r.State(eventType, stateKey); err == nil {
		j, err := json.Marshal(state)
		if err != nil {
			return nil, fmt.Errorf("json.Marshal: %w", err)
		}
		return NewSourceFromJSON(j)
	} else {
		return nil, err
	}
}
