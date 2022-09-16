package mapshot

import (
	"fmt"
	"time"

	"github.com/vikpe/go-ezquake"
)

type IClientController interface {
	Command(cmd string)
	CommandWithOptions(cmd string, options ezquake.CommandOptions)
}

type Client struct {
	controller IClientController
}

func NewClient(controller IClientController) *Client {
	client := Client{
		controller: controller,
	}

	client.do(`hide all; crosshair 0; sshot_autoname 1"`, 50*time.Millisecond)
	return &client
}

func (c *Client) Mapshot(mapName, settings string) {
	c.do(fmt.Sprintf("map %s", mapName), 3*time.Second)
	c.do(settings, 50*time.Millisecond)
	c.do("clear; wait; screenshot", 500*time.Millisecond)
}

func (c *Client) do(cmd string, timeout time.Duration) {
	c.controller.CommandWithOptions(cmd, ezquake.CommandOptions{Timeout: timeout})
}
