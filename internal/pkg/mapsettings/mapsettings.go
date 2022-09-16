package mapsettings

import (
	"encoding/json"
	"io/ioutil"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

type MapSettings map[string]string

func (s MapSettings) HasMap(mapName string) bool {
	_, hasKey := s[mapName]
	return hasKey
}

func (s MapSettings) MapNames() []string {
	mapNames := maps.Keys(s)
	slices.Sort(mapNames)
	return mapNames
}

func FromJsonFile(filePath string) (MapSettings, error) {
	file, err := ioutil.ReadFile(filePath)

	if err != nil {
		return MapSettings{}, err
	}

	settings := MapSettings{}
	err = json.Unmarshal(file, &settings)

	return settings, err
}
