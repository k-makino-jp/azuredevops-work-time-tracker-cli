/*
Copyright © 2022 k-makino-jp

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmdController represents the base command when called without any subcommands.
var rootCmdController = &cobra.Command{
	Use:   "devopsctl",
	Short: "devopsctl manages the azure devops work items",
	Long:  "devopsctl manages the azure devops work items",
}

// GetRootCmdController returns rootCmdController instance.
func GetRootCmdController() *cobra.Command {
	return rootCmdController
}
