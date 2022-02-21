package data_access

import (
	"github.com/k-makino-jp/azuredevops-work-time-tracker-cli/entity"
	"github.com/k-makino-jp/azuredevops-work-time-tracker-cli/infrastructure/azuredevops"
	"github.com/k-makino-jp/azuredevops-work-time-tracker-cli/usecase"
	"github.com/microsoft/azure-devops-go-api/azuredevops/workitemtracking"
)

type AzureDevOpsClient interface {
	GetUpdates(usecase.GetCmdOptions) (*[]workitemtracking.WorkItemUpdate, error)
}

type workItemUpdateDataAccessor struct {
	client AzureDevOpsClient
}

func NewWorkItemUpdateDataAccessor(config entity.Config) *workItemUpdateDataAccessor {
	return &workItemUpdateDataAccessor{
		client: azuredevops.NewClient(config),
	}
}

func (w workItemUpdateDataAccessor) Get(getCmdOptions usecase.GetCmdOptions) (*[]workitemtracking.WorkItemUpdate, error) {
	return w.client.GetUpdates(getCmdOptions)
}
