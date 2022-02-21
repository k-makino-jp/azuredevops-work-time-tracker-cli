package repository

import (
	"github.com/k-makino-jp/azuredevops-work-time-tracker-cli/entity"
	"github.com/k-makino-jp/azuredevops-work-time-tracker-cli/infrastructure/azuredevops"
	"github.com/k-makino-jp/azuredevops-work-time-tracker-cli/usecase"
	"github.com/microsoft/azure-devops-go-api/azuredevops/workitemtracking"
)

type AzureDevOpsClient interface {
	GetUpdates(usecase.GetCmdOptions) (*[]workitemtracking.WorkItemUpdate, error)
}

type workItemUpdateRepository struct {
	client AzureDevOpsClient
}

func NewWorkItemUpdateRepository(config entity.Config) *workItemUpdateRepository {
	return &workItemUpdateRepository{
		client: azuredevops.NewClient(config),
	}
}

func (w workItemUpdateRepository) Get(getCmdOptions usecase.GetCmdOptions) (*[]workitemtracking.WorkItemUpdate, error) {
	return w.client.GetUpdates(getCmdOptions)
}
