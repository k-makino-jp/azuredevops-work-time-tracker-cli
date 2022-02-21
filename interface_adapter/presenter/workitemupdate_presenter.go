package presenter

import (
	"fmt"
	"time"

	"github.com/k-makino-jp/azuredevops-work-time-tracker-cli/usecase"
	"github.com/microsoft/azure-devops-go-api/azuredevops/workitemtracking"
)

var (
	MaxTime time.Time = time.Date(9999, 1, 1, 0, 0, 0, 0, time.UTC)
)

type WorkItemUpdatePresenter interface{}

type workItemUpdatePresenter struct{}

func NewWorkItemUpdatePresenter() *workItemUpdatePresenter {
	return &workItemUpdatePresenter{}
}

func (w workItemUpdatePresenter) Calculate(targetState usecase.State, updateList *[]workitemtracking.WorkItemUpdate) {
	fmt.Printf("# Target State: %s\n\n", targetState.String())
	fmt.Printf("## Histories\n\n")
	fmt.Println("| Start Time           | End Time             | Spend Time [hours]   |")
	fmt.Println("| -------------------- | ---------------------| -------------------- |")
	var totalDiff time.Duration
	for _, update := range *updateList {
		if (*update.Fields)["System.State"].NewValue == targetState.String() {
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
			formattedStartTime := startTime.Format("2006-01-02T15:04:05Z")
			formattedEndTime := endTime.Format("2006-01-02T15:04:05Z")
			fmt.Printf("| %s | %s | %05.2f                |\n", formattedStartTime, formattedEndTime, spendTime.Hours())
		}
	}
	fmt.Println("")
	fmt.Printf("## Total Spend Time\n\n")
	fmt.Printf("* %05.2f [hours]\n\n", totalDiff.Hours())
}
