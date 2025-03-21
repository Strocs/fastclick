package settings

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Settings represents the application configuration
type Settings struct {
	TriggerKey string `json:"trigger_key"`
}

func getConfigFilePath() string {
	return filepath.Join(".", "fastClickerConfig.json")
}

func LoadSettings() (Settings, error) {

	defaultSettings := Settings{
		TriggerKey: "Enter",
	}

	data, err := os.ReadFile(getConfigFilePath())

	if err != nil {
		if os.IsNotExist(err) {
			SaveSettings(defaultSettings)
			return defaultSettings, nil
		}
	}

	var settings Settings
	err = json.Unmarshal(data, &settings)

	if err != nil {
		panic(err)
	}

	return settings, nil
}

func SaveSettings(settings Settings) error {

	data, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(getConfigFilePath(), data, 0666) // 0644 is the file permission
	if err != nil {
		panic(err)
	}

	return nil
}

func UpdateSettings(newTriggerKey string) (Settings, error) {

	settings, err := LoadSettings()
	if err != nil {
		panic(err)
	}

	settings.TriggerKey = newTriggerKey
	SaveSettings(settings)
	return settings, nil
}
