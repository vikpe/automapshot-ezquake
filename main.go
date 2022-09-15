package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/vikpe/go-ezquake"
)

type MapSettings map[string]string

func main() {
	// load env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("unable to load environment variables", err)
		return
	}

	// take screenshots
	ctrl := ezquake.NewClientController(
		os.Getenv("EZQUAKE_PROCESS_USERNAME"),
		os.Getenv("EZQUAKE_BIN_PATH"),
	)

	if !ctrl.Process.IsStarted() {
		log.Fatal("ezQuake is not started")
		return
	}

	mapSettings := getMapSettingsFromFile("mapsettings.json")

	doCmd := func(cmd string, timeout time.Duration) {
		ctrl.CommandWithOptions(cmd, ezquake.CommandOptions{Timeout: timeout})
	}

	for mapName, settings := range mapSettings {
		fmt.Println(mapName, settings)
		doCmd(fmt.Sprintf("map %s", mapName), time.Second*4)
		doCmd(settings, time.Millisecond*250)
		doCmd("clear", time.Millisecond*250)
		doCmd("screenshot", time.Second)
	}

	// todo: move screenshots to /dist
	// todo: resize screenshots
}

func getMapSettingsFromFile(filePath string) MapSettings {
	file, _ := ioutil.ReadFile(filePath)
	settings := MapSettings{}
	_ = json.Unmarshal(file, &settings)
	return settings
}
