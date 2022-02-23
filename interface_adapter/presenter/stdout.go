package presenter

import (
	"fmt"

	"github.com/k-makino-jp/azuredevops-work-time-tracker-cli/usecase"
)

type workItemUpdateStdoutPresenter struct{}

func NewWorkItemUpdateStdoutPresenter() *workItemUpdateStdoutPresenter {
	return &workItemUpdateStdoutPresenter{}
}

func (w workItemUpdateStdoutPresenter) Output(workItemUpdateHistories usecase.WorkItemUpdateHistories) {
	fmt.Printf("# Target State: %s\n\n", workItemUpdateHistories.TargetState.String())
	fmt.Printf("## Histories\n\n")
	fmt.Println("| Start Time           | End Time             | Spend Time [hours]   |")
	fmt.Println("| -------------------- | -------------------- | -------------------- |")
	for _, workItemUpdateHistory := range workItemUpdateHistories.Histories {
		formattedStartTime := workItemUpdateHistory.StartTime.Format("2006-01-02T15:04:05Z")
		formattedEndTime := workItemUpdateHistory.EndTime.Format("2006-01-02T15:04:05Z")
		fmt.Printf("| %s | %s | %05.2f                |\n", formattedStartTime, formattedEndTime, workItemUpdateHistory.SpendTime.Hours())
	}
	fmt.Println("")
	fmt.Printf("## Total Spend Time\n\n")
	fmt.Printf("* %05.2f [hours]\n\n", workItemUpdateHistories.TotalSpendTime.Hours())
}
