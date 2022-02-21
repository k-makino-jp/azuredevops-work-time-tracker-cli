package config

import (
	"encoding/json"
	"os"

	"github.com/k-makino-jp/azuredevops-work-time-tracker-cli/entity"
)

type reader struct{}

func NewReader() *reader {
	return &reader{}
}

func (r reader) Read() (entity.Config, error) {
	bytes, err := os.ReadFile("config.json")
	if err != nil {
		return entity.Config{}, err
	}

	var config entity.Config
	err = json.Unmarshal(bytes, &config)
	return config, err
}
