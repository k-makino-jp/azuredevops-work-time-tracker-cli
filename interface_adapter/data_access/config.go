package data_access

import (
	"encoding/json"
	"os"

	"github.com/k-makino-jp/azuredevops-work-time-tracker-cli/entity"
)

const (
	configFilePath = "config.json"
)

type configDataAccessor struct{}

func NewConfigDataAccessor() *configDataAccessor {
	return &configDataAccessor{}
}

func (c configDataAccessor) Read() (entity.Config, error) {
	bytes, err := os.ReadFile(configFilePath)
	if err != nil {
		return entity.Config{}, err
	}

	var config entity.Config
	err = json.Unmarshal(bytes, &config)
	return config, err
}
