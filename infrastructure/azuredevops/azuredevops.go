package azuredevops

import (
	"context"
	"net/url"

	"github.com/k-makino-jp/azuredevops-work-time-tracker-cli/entity"
	"github.com/k-makino-jp/azuredevops-work-time-tracker-cli/usecase"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/workitemtracking"
)

type client struct {
	Url     *url.URL
	Id      int
	Project string
}

const (
	HttpsSchema     = "https"
	AzureDevOpsHost = "dev.azure.com"
)

func NewClient(config entity.Config) *client {
	return &client{
		Url: &url.URL{
			Scheme: HttpsSchema,
			Host:   AzureDevOpsHost,
			Path:   config.Organization,
		},
		Id:      config.Id,
		Project: config.Project,
	}
}

func (a client) GetUpdates(getCmdOptions usecase.GetCmdOptions) (*[]workitemtracking.WorkItemUpdate, error) {
	connection := azuredevops.NewPatConnection(a.Url.String(), getCmdOptions.Pat)
	ctx := context.Background()

	workItemTrackingClient, err := workitemtracking.NewClient(ctx, connection)
	if err != nil {
		return nil, err
	}

	updatesArgs := workitemtracking.GetUpdatesArgs{
		Id:      &a.Id,
		Project: &a.Project,
	}
	return workItemTrackingClient.GetUpdates(ctx, updatesArgs)
}
