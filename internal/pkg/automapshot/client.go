package automapshot

import (
	"errors"
	"fmt"
	"time"

	"github.com/vikpe/go-ezquake"
)

type Client struct {
	controller *ezquake.ClientController
}

func NewClient(username, binPath string) *Client {
	client := Client{
		controller: ezquake.NewClientController(username, binPath),
	}

	client.do(`hide all; crosshair 0; sshot_autoname 1; sshot_dir mapshots"`, 50*time.Millisecond)
	return &client
}

func (c *Client) Mapshots(mapNames []string, settings MapSettings) error {
	if len(mapNames) == 1 && "all" == mapNames[0] {
		mapNames = settings.MapNames()
	}

	for _, mapName := range mapNames {
		if !settings.HasMap(mapName) {
			continue
		}

		err := c.Mapshot(mapName, settings[mapName])

		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) Mapshot(mapName, mapSettings string) error {
	if !c.controller.Process.IsStarted() {
		return errors.New("ezquake is not started")
	}

	c.do(fmt.Sprintf("map %s", mapName), 3*time.Second)
	c.do(mapSettings, 50*time.Millisecond)
	c.do("clear; wait; screenshot", 500*time.Millisecond)

	return nil
}

func (c *Client) do(cmd string, timeout time.Duration) {
	c.controller.CommandWithOptions(cmd, ezquake.CommandOptions{Timeout: timeout})
}
