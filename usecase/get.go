package usecase

import (
	"fmt"
	"strings"
	"time"

	"github.com/microsoft/azure-devops-go-api/azuredevops/workitemtracking"
)

type State string

func (s State) String() string {
	return string(s)
}

func (s State) IsValid() bool {
	switch s {
	case StateNew:
		return true
	case StateActive:
		return true
	case StateClosed:
		return true
	}
	return false
}

const (
	StateNew    State = "New"
	StateActive State = "Active"
	StateClosed State = "Closed"
)

type GetCmdOptions struct {
	Id     int
	Pat    string
	States string
}

type WorkItemUpdateHistories struct {
	TargetState    State
	Histories      []WorkItemUpdateHistory
	TotalSpendTime time.Duration
}

type WorkItemUpdateHistory struct {
	StartTime time.Time
	EndTime   time.Time
	SpendTime time.Duration
}

// GetCmdInteractor implements Cmd interface.
type GetCmdInteractor struct {
	GetCmdOptions            *GetCmdOptions
	WorkItemUpdateTracker    WorkItemUpdateTracker
	WorkItemUpdateCalculator WorkItemUpdateCalculator
	WorkItemUpdatePresenter  WorkItemUpdatePresenter
}

type WorkItemUpdateTracker interface {
	Get(id int, pat string) (*[]workitemtracking.WorkItemUpdate, error)
}

type WorkItemUpdateCalculator interface {
	Calculate(targetState State, updateList *[]workitemtracking.WorkItemUpdate) WorkItemUpdateHistories
}

type WorkItemUpdatePresenter interface {
	Output(WorkItemUpdateHistories)
}

func (g GetCmdInteractor) Execute() error {
	workItemUpdates, err := g.WorkItemUpdateTracker.Get(g.GetCmdOptions.Id, g.GetCmdOptions.Pat)
	if err != nil {
		return err
	}

	states := strings.Split(g.GetCmdOptions.States, ",")
	for _, state := range states {
		if !State(state).IsValid() {
			fmt.Printf("WARNING: Invalid work item state %s is specified. Available states are \"New\", \"Active\" and \"Closed\"\n", state)
			continue
		}
		workItemUpdateHistories := g.WorkItemUpdateCalculator.Calculate(State(state), workItemUpdates)
		g.WorkItemUpdatePresenter.Output(workItemUpdateHistories)
	}
	return nil
}
