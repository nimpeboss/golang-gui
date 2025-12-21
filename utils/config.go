package utils

import (
	"encoding/json"
	"os"

	"advanced-gui/state"
)

func SaveConfig(s *state.AppState) error {
	file, err := os.Create("config.json")
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(s)
}

func LoadConfig(s *state.AppState) error {
	file, err := os.Open("config.json")
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(s)
}
