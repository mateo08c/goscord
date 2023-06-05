package goscord

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/gateway"
	"github.com/Goscord/goscord/goscord/gateway/event"
	"github.com/Goscord/goscord/goscord/rest"
)

type Client struct {
	Session *gateway.Session
	Rest    *rest.Client
}

// New creates a new Client instance.
func New(options *gateway.Options) *Client {
	c := new(Client)
	//Create a new rest Client for rest requests
	c.Rest = rest.NewClient(options.Token)

	//Create a new gateway session for events
	c.Session = gateway.NewSession(options, c.Rest)

	return c
}

// On registers an event handler.
func (c *Client) On(event event.EventType, fun any) error {
	return c.Session.On(event, fun)
}

// Me returns the user of the Client.
func (c *Client) Me() *discord.User {
	return c.Session.Me()
}

// Login logs the websocket Client in.
func (c *Client) Login() error {
	return c.Session.Login()
}

// Close closes the websocket connection.
func (c *Client) Close() {
	c.Session.Close()
}

func (c *Client) GetGuild(s string) (*Guild, error) {
	guild, err := c.Session.State().Guild(s)
	if err != nil {
		get, err := c.Rest.Guild.Get(discord.Snowflake(s))
		if err != nil {
			return nil, err
		}

		return &Guild{Guild: get, Client: c}, nil
	}

	return &Guild{Guild: guild, Client: c}, nil
}
