package todo

import "time"

type (
	Task struct {
		TaskId       int       `json:"task_id"`
		PersonName   string    `json:"person_name"`
		TaskName     string    `json:"task_name"`
		DeadlineDate time.Time `json:"deadline_date"`
		TaskStatus   string    `json:"task_dtatus"`
	}

	TaskResponse struct {
		Tasks []Task `json:"tasks"`
	}

	TaskDetailResponse struct {
		Task Task `json:"task"`
	}

	TaskDeleteResponse struct {
	}
)
