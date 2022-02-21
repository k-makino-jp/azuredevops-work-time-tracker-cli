/*
Copyright © 2022 k-makino-jp

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rootCmdController represents the base command when called without any subcommands.
var rootCmdController = &cobra.Command{
	Use:   "azuredevops-work-time-tracker-cli",
	Short: "short description of root subcommand",
	Long:  `long description of root subcommand`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rootCmd called.")
	},
}

func GetRootCmdController() *cobra.Command {
	return rootCmdController
}
