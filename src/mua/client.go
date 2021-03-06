package mua

import (
	"fmt"

	"github.com/matrix-org/gomatrix"
)

type Client struct {
	client *gomatrix.Client
	state  *Lua
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
			InitialDeviceDisplayName: "Mua",
		})
		if err != nil {
			return nil, fmt.Errorf("client.RegisterGuest: %w", err)
		}
		client.UserID = register.UserID
		client.AccessToken = register.AccessToken
	}

	c := &Client{
		client: client,
		rooms:  make(map[string]*Room),
	}

	vm, err := NewLua(c)
	if err != nil {
		return nil, fmt.Errorf("NewLua: %w", err)
	}
	c.state = vm

	return c, nil
}

func (c *Client) UserID() string {
	return c.client.UserID
}

func (c *Client) AccessToken() string {
	return c.client.AccessToken
}

func (c *Client) Execute(source string) error {
	return c.state.Execute(source)
}

func (c *Client) ExecuteFile(file string) error {
	return c.state.ExecuteFile(file)
}
