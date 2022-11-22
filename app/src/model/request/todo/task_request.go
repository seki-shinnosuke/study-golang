package todo

import "time"

type (
	UriParam struct {
		Id int `uri:"id" binding:"required"`
	}

	Task struct {
		PersonName   string    `json:"person_name"`
		TaskName     string    `json:"task_name"`
		DeadlineDate time.Time `json:"deadline_date"`
		TaskStatus   string    `json:"task_dtatus"`
	}
)
