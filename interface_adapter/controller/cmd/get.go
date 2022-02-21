/*
Copyright © 2022 k-makino-jp

*/
package cmd

import (
	"log"

	"github.com/k-makino-jp/azuredevops-work-time-tracker-cli/interface_adapter/data_access"
	"github.com/k-makino-jp/azuredevops-work-time-tracker-cli/interface_adapter/presenter"
	"github.com/k-makino-jp/azuredevops-work-time-tracker-cli/usecase"
	"github.com/spf13/cobra"
)

var (
	useCaseGetCmd usecase.Cmd
	pat           string

	// getCmd represents the get command
	getCmd = &cobra.Command{
		Use:   "get",
		Short: "get short description",
		Long:  "get long description",
		RunE: func(cmd *cobra.Command, args []string) error {
			return useCaseGetCmd.Execute()
		},
	}
)

func init() {
	getCmdOptions := usecase.GetCmdOptions{}
	getCmd.Flags().StringVarP(&getCmdOptions.Pat, "pat", "p", "", "Azure DevOps Personal Access Token (required)")
	getCmd.MarkFlagRequired("pat")
	rootCmdController.AddCommand(getCmd)

	configRepository := data_access.NewConfigDataAccessor()
	config, err := configRepository.Read()
	if err != nil {
		log.Fatal(err)
	}

	useCaseGetCmd = usecase.GetCmdInteractor{
		GetCmdOptions:            &getCmdOptions,
		WorkItemUpdateTracker:    data_access.NewWorkItemUpdateDataAccessor(config),
		WorkItemUpdateCalculator: presenter.NewWorkItemUpdateCalculator(),
		WorkItemUpdatePresenter:  presenter.NewWorkItemUpdateStdoutPresenter(),
	}
}
