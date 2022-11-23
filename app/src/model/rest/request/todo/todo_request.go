package todo

import (
	"github.com/volatiletech/null/v8"

	code "github.com/seki-shinnosuke/study-golang/code"
)

type (
	UriParam struct {
		TaskId int `uri:"taskId" binding:"required,min=1"`
	}

	Task struct {
		PersonName   string    `json:"person_name" binding:"required,max=128"`
		TaskName     string    `json:"task_name" binding:"required,max=128"`
		DeadlineDate null.Time `json:"deadline_date"`
		TaskStatus   string    `json:"task_dtatus" binding:"required"`
	}
)

func (task *Task) Validate() bool {
	// Check TaskStatus
	return code.NameToTaskStatusCode(task.TaskStatus) != nil
}
