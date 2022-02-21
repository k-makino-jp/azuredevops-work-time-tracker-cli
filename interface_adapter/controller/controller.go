package controller

import "github.com/k-makino-jp/azuredevops-work-time-tracker-cli/interface_adapter/controller/cmd"

// Controller defines the required methods for executing the use case.
type Controller interface {
	Execute() error
}

// GetController returns a controller for handling user inputs.
func GetController() Controller {
	return cmd.GetRootCmdController()
}
