package presenter

import (
	"time"

	"github.com/k-makino-jp/azuredevops-work-time-tracker-cli/usecase"
	"github.com/microsoft/azure-devops-go-api/azuredevops/workitemtracking"
)

var (
	MaxTime time.Time = time.Date(9999, 1, 1, 0, 0, 0, 0, time.UTC)
)

type workItemUpdateCalculator struct{}

func NewWorkItemUpdateCalculator() *workItemUpdateCalculator {
	return &workItemUpdateCalculator{}
}

func (w workItemUpdateCalculator) Calculate(targetState usecase.State, updateList *[]workitemtracking.WorkItemUpdate) usecase.WorkItemUpdateHistories {
	var workItemUpdateHistories usecase.WorkItemUpdateHistories
	var totalDiff time.Duration
	for _, update := range *updateList {
		if (*update.Fields)["System.State"].NewValue != targetState.String() {
			continue
		}

		startTimeStr, ok := (*update.Fields)["System.RevisedDate"].OldValue.(string)
		if !ok {
			startTimeStr = (*update.Fields)["System.CreatedDate"].NewValue.(string)
		}
		startTime, _ := time.Parse("2006-01-02T15:04:05Z", startTimeStr)

		endTimeStr := (*update.Fields)["System.RevisedDate"].NewValue.(string)
		endTime, _ := time.Parse("2006-01-02T15:04:05Z", endTimeStr)
		if endTime == MaxTime {
			endTime = time.Now().UTC()
		}

		spendTime := endTime.Sub(startTime)
		totalDiff += spendTime

		workItemUpdateHistory := usecase.WorkItemUpdateHistory{
			StartTime: startTime,
			EndTime:   endTime,
			SpendTime: spendTime,
		}
		workItemUpdateHistories.Histories = append(workItemUpdateHistories.Histories, workItemUpdateHistory)
	}
	workItemUpdateHistories.TargetState = targetState
	workItemUpdateHistories.TotalSpendTime = totalDiff
	return workItemUpdateHistories
}
