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

	for mapName, settings := range mapSettings {
		fmt.Println(mapName, settings)
		ctrl.CommandWithOptions(fmt.Sprintf("map %s", mapName), ezquake.CommandOptions{Timeout: time.Second * 4})
		ctrl.CommandWithOptions(settings, ezquake.CommandOptions{Timeout: time.Millisecond * 250})
		ctrl.CommandWithOptions("clear", ezquake.CommandOptions{Timeout: time.Millisecond * 250})
		ctrl.CommandWithOptions("screenshot", ezquake.CommandOptions{Timeout: time.Second})
	}

	// todo: collect screenshots
	// todo: resize screenshots
}

func getMapSettingsFromFile(filePath string) MapSettings {
	file, _ := ioutil.ReadFile(filePath)
	settings := MapSettings{}
	_ = json.Unmarshal(file, &settings)
	return settings
}
