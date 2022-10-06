package mapshot

import (
	"errors"
	"fmt"
	"path/filepath"
	"time"

	"github.com/vikpe/go-ezquake"
)

type IClientController interface {
	Command(cmd string, options ezquake.CommandOptions)
	CommandWithOptions(cmd string, options ezquake.CommandOptions)
}

type Client struct {
	controller *ezquake.ClientController
}

func NewClient(username, binPath string) *Client {
	client := Client{
		controller: ezquake.NewClientController(username, binPath),
	}

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

	err := c.loadMap(mapName)

	if err != nil {
		return err
	}

	c.do(mapSettings, 50*time.Millisecond)
	c.do("clear; wait; screenshot", 500*time.Millisecond)

	return nil
}

func (c *Client) loadMap(mapName string) error {
	assets := ezquake.NewAssetManager(filepath.Dir(c.controller.Process.Path))

	if !assets.HasMap(mapName) {
		err := assets.DownloadMap(mapName)

		if err != nil {
			return errors.New(fmt.Sprintf("%s was not found in qw/maps and could not be downloaded (%s)", mapName, ezquake.MapUrl(mapName)))
		}

		time.Sleep(time.Millisecond * 50)
	}

	c.do(fmt.Sprintf("map %s", mapName), 3*time.Second)
	return nil
}

func (c *Client) do(cmd string, timeout time.Duration) {
	c.controller.CommandWithOptions(cmd, ezquake.CommandOptions{Delay: timeout})
}
