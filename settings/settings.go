package settings

import (
	"encoding/json"
	"os"
)

type Settings struct {
	IpUrl           string `json:"ip_url"`
	LogFile         string `json:"log_filename"`
	DataFile        string `json:"data_filename"`
	BotUrl          string `json:"bot_url"`
	ChatId          string `json:"chat_id"`
	TelegramMessage string `json:"telegram_message"`
}

func ReadFileSettings(path string) (*Settings, error) {

	if path == "nil" {
		path = "settings.json"
	}

	var settings *Settings

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&settings)
	if err != nil {
		return nil, err
	}

	return settings, nil
}
