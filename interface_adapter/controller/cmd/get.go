/*
Copyright © 2022 k-makino-jp

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/k-makino-jp/azuredevops-work-time-tracker-cli/interface_adapter/data_access"
	"github.com/k-makino-jp/azuredevops-work-time-tracker-cli/interface_adapter/presenter"
	"github.com/k-makino-jp/azuredevops-work-time-tracker-cli/usecase"
	"github.com/spf13/cobra"
)

var (
	useCaseGetCmd usecase.Cmd

	// getCmd represents the get command
	getCmd = &cobra.Command{
		Use:     "get",
		Short:   "Get work item updates histories.",
		Long:    "Get work item updates histories.",
		Example: `devopsctl get -i 100 -p "personalaccesstoken -s "New,Active""`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := useCaseGetCmd.Execute(); err != nil {
				fmt.Println("ERROR:", err)
			}
		},
	}
)

func init() {
	getCmdOptions := usecase.GetCmdOptions{}
	getCmd.Flags().IntVarP(&getCmdOptions.Id, "id", "i", 0, "Azure DevOps WorkItem ID (required)")
	getCmd.Flags().StringVarP(&getCmdOptions.Pat, "pat", "p", "", "Azure DevOps Personal Access Token (required)")
	getCmd.Flags().StringVarP(&getCmdOptions.Statuses, "states", "s", "Active", "Azure DevOps WorkItem Statuses (comma separated values)")
	getCmd.MarkFlagRequired("id")
	getCmd.MarkFlagRequired("pat")
	rootCmdController.AddCommand(getCmd)

	configRepository := data_access.NewConfigDataAccessor()
	config, err := configRepository.Read()
	if err != nil {
		log.Fatalln("ERROR:", err)
	}

	useCaseGetCmd = usecase.GetCmdInteractor{
		GetCmdOptions:            &getCmdOptions,
		WorkItemUpdateTracker:    data_access.NewWorkItemUpdateDataAccessor(config),
		WorkItemUpdateCalculator: presenter.NewWorkItemUpdateCalculator(),
		WorkItemUpdatePresenter:  presenter.NewWorkItemUpdateStdoutPresenter(),
	}
}
