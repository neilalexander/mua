package mdl

import (
	"fmt"

	"github.com/matrix-org/gomatrix"
)

type Client struct {
	client *gomatrix.Client
	rooms  map[string]*Room
}

func NewClient(hsURL, userID, accessToken string) (*Client, error) {
	fmt.Println("Using", hsURL)

	client, err := gomatrix.NewClient(hsURL, userID, accessToken)
	if err != nil {
		return nil, fmt.Errorf("gomatrix.NewClient: %w", err)
	}

	if userID == "" || accessToken == "" {
		fmt.Println("Registering guest user")
		register, _, err := client.RegisterGuest(&gomatrix.ReqRegister{
			InitialDeviceDisplayName: "MLD",
		})
		if err != nil {
			return nil, fmt.Errorf("client.RegisterGuest: %w", err)
		}
		client.UserID = register.UserID
		client.AccessToken = register.AccessToken
	}

	return &Client{
		client: client,
		rooms:  make(map[string]*Room),
	}, nil
}

func (c *Client) UserID() string {
	return c.client.UserID
}

func (c *Client) AccessToken() string {
	return c.client.AccessToken
}

func (c *Client) NewRoom(roomID string) (*Room, error) {
	var err error
	/*
		var roomID string
		resp, err := c.client.JoinRoom(roomIDOrAlias, "", map[string]string{
			"display_name": "MLD",
		})

		switch e := err.(type) {
		case nil:
			roomID = resp.RoomID
		case gomatrix.HTTPError:
			if e.Code == 404 {
				create, createErr := c.client.CreateRoom(&gomatrix.ReqCreateRoom{
					Name:       "MLD Room",
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

	room.state, err = room.NewLua()
	if err != nil {
		return nil, fmt.Errorf("room.NewLua: %w", err)
	}

	return room, nil
}
