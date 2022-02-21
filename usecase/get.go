package usecase

import (
	"fmt"
	"time"

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
	GetCmdOptions            *GetCmdOptions
	WorkItemUpdateTracker    WorkItemUpdateTracker
	WorkItemUpdateCalculator WorkItemUpdateCalculator
	WorkItemUpdatePresenter  WorkItemUpdatePresenter
}

type WorkItemUpdateTracker interface {
	Get(GetCmdOptions) (*[]workitemtracking.WorkItemUpdate, error)
}

type WorkItemUpdateHistories struct {
	TargetState    State
	HistoryList    []WorkItemUpdateHistory
	TotalSpendTime time.Duration
}

type WorkItemUpdateHistory struct {
	StartTime time.Time
	EndTime   time.Time
	SpendTime time.Duration
}

type WorkItemUpdateCalculator interface {
	Calculate(targetState State, updateList *[]workitemtracking.WorkItemUpdate) WorkItemUpdateHistories
}

type WorkItemUpdatePresenter interface {
	Output(WorkItemUpdateHistories)
}

func (g GetCmdInteractor) Execute() error {
	workItemUpdates, err := g.WorkItemUpdateTracker.Get(*g.GetCmdOptions)
	if err != nil {
		fmt.Println(err)
		return err
	}

	newWorkItemUpdateHistories := g.WorkItemUpdateCalculator.Calculate(StateNew, workItemUpdates)
	g.WorkItemUpdatePresenter.Output(newWorkItemUpdateHistories)

	activeWorkItemUpdateHistories := g.WorkItemUpdateCalculator.Calculate(StateActive, workItemUpdates)
	g.WorkItemUpdatePresenter.Output(activeWorkItemUpdateHistories)
	return nil
}
