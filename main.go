/*
Copyright © 2022 k-makino-jp

*/
package main

import (
	"os"

	"github.com/k-makino-jp/azuredevops-work-time-tracker-cli/interface_adapter/controller"
)

func main() {
	const successfulExitCode, failureExitCode int = 0, 1
	controller := controller.GetController()
	if err := controller.Execute(); err != nil {
		os.Exit(failureExitCode)
	}
	os.Exit(successfulExitCode)
}
