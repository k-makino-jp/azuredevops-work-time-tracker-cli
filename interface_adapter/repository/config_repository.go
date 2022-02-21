package repository

import (
	"github.com/k-makino-jp/azuredevops-work-time-tracker-cli/entity"
	"github.com/k-makino-jp/azuredevops-work-time-tracker-cli/infrastructure/config"
)

type ConfigReader interface {
	Read() (entity.Config, error)
}

type configRepository struct {
	reader ConfigReader
}

func NewConfigRepository() *configRepository {
	return &configRepository{
		reader: config.NewReader(),
	}
}

func (c configRepository) Read() (entity.Config, error) {
	return c.reader.Read()
}
