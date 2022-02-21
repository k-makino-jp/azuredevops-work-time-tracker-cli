package usecase

import (
	"fmt"

	"github.com/microsoft/azure-devops-go-api/azuredevops/workitemtracking"
)

type State string

func (s State) String() string {
	return string(s)
}

const (
	StateNew    State = "New"
	StateActive State = "Active"
	StateClosed State = "Closed"
)

type GetCmdOptions struct {
	Pat string
}

// GetCmdInteractor implements Cmd interface.
type GetCmdInteractor struct {
	GetCmdOptions               *GetCmdOptions
	WorkItemUpdateTracker       WorkItemUpdateTracker
	WorkItemSpendTimeCalculator WorkItemSpendTimeCalculator
}

type WorkItemUpdateTracker interface {
	Get(GetCmdOptions) (*[]workitemtracking.WorkItemUpdate, error)
}

type WorkItemSpendTimeCalculator interface {
	Calculate(State, *[]workitemtracking.WorkItemUpdate)
}

func (g GetCmdInteractor) Execute() error {
	workItemUpdates, err := g.WorkItemUpdateTracker.Get(*g.GetCmdOptions)
	if err != nil {
		fmt.Println(err)
		return err
	}
	g.WorkItemSpendTimeCalculator.Calculate(StateNew, workItemUpdates)
	g.WorkItemSpendTimeCalculator.Calculate(StateActive, workItemUpdates)
	return nil
}
