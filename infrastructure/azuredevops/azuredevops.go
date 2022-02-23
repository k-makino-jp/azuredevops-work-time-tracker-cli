package azuredevops

import (
	"context"
	"net/url"

	"github.com/k-makino-jp/azuredevops-work-time-tracker-cli/entity"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/workitemtracking"
)

const (
	httpsSchema     = "https"
	azureDevOpsHost = "dev.azure.com"
)

type client struct {
	Url     *url.URL
	Id      int
	Project string
}

// NewClient returns instance of azure devops rest client.
func NewClient(config entity.Config) *client {
	return &client{
		Url: &url.URL{
			Scheme: httpsSchema,
			Host:   azureDevOpsHost,
			Path:   config.Organization,
		},
		Project: config.Project,
	}
}

// GetUpdates calls "Updates - List" API.
func (a client) GetUpdates(id int, pat string) (*[]workitemtracking.WorkItemUpdate, error) {
	connection := azuredevops.NewPatConnection(a.Url.String(), pat)
	ctx := context.Background()

	workItemTrackingClient, err := workitemtracking.NewClient(ctx, connection)
	if err != nil {
		return nil, err
	}

	updatesArgs := workitemtracking.GetUpdatesArgs{
		Id:      &id,
		Project: &a.Project,
	}
	return workItemTrackingClient.GetUpdates(ctx, updatesArgs)
}
