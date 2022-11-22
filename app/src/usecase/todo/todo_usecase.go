package todo

import (
	"context"

	models "github.com/seki-shinnosuke/study-golang/model/db"
	response "github.com/seki-shinnosuke/study-golang/model/response/todo"
)

type TodoUsecase struct {
}

func NewTodoUsecase() *TodoUsecase {
	return &TodoUsecase{}
}

func (uc *TodoUsecase) GetTodos() (*response.TaskResponse, error) {
	ctx := context.Background()
	todos, err := models.TaskManagements().AllG(ctx)

	if err != nil {
		return nil, err
	}

	tasks := make([]response.Task, 0, len(todos))
	for _, todo := range todos {
		tasks = append(tasks, response.Task{
			TaskId:       uint64(todo.TaskID),
			PersonName:   todo.PersonName,
			TaskName:     todo.TaskName,
			DeadlineDate: todo.DeadlineDate.Time,
			TaskStatus:   todo.TaskStatus,
		})
	}

	return &response.TaskResponse{Tasks: tasks}, nil
}
