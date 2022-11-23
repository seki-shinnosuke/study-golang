package todo

import (
	"github.com/volatiletech/null/v8"
)

type (
	Task struct {
		TaskId       int       `json:"task_id"`
		PersonName   string    `json:"person_name"`
		TaskName     string    `json:"task_name"`
		DeadlineDate null.Time `json:"deadline_date"`
		TaskStatus   string    `json:"task_dtatus"`
	}

	TodoResponse struct {
		Tasks []Task `json:"tasks"`
	}

	TodoDetailResponse struct {
		Task Task `json:"task"`
	}

	TodoDeleteResponse struct {
	}
)
