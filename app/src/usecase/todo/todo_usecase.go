package todo

import (
	"context"
	"database/sql"

	e "github.com/seki-shinnosuke/study-golang/error"
	model "github.com/seki-shinnosuke/study-golang/model/db"
	response "github.com/seki-shinnosuke/study-golang/model/response/todo"
)

type TodoUsecase struct {
}

func NewTodoUsecase() *TodoUsecase {
	return &TodoUsecase{}
}

func (uc *TodoUsecase) GetTodos() (*response.TaskResponse, error) {
	ctx := context.Background()
	todos, err := model.TaskManagements().AllG(ctx)

	if err != nil {
		return nil, e.WithError(err, e.InternalServerError)
	}

	tasks := make([]response.Task, 0, len(todos))
	for _, todo := range todos {
		tasks = append(tasks, response.Task{
			TaskId:       todo.TaskID,
			PersonName:   todo.PersonName,
			TaskName:     todo.TaskName,
			DeadlineDate: todo.DeadlineDate.Time,
			TaskStatus:   todo.TaskStatus,
		})
	}

	return &response.TaskResponse{Tasks: tasks}, nil
}

func (uc *TodoUsecase) GetTodo(targetTaskId int) (*response.TaskDetailResponse, error) {
	ctx := context.Background()
	todo, err := model.FindTaskManagementG(ctx, targetTaskId)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, e.WithError(err, e.DataNotFound)
		}
		return nil, e.WithError(err, e.InternalServerError)
	}

	return &response.TaskDetailResponse{Task: response.Task{
		TaskId:       todo.TaskID,
		PersonName:   todo.PersonName,
		TaskName:     todo.TaskName,
		DeadlineDate: todo.DeadlineDate.Time,
		TaskStatus:   todo.TaskStatus,
	}}, nil
}

func (uc *TodoUsecase) DeleteTodo(targetTaskId int) error {
	ctx := context.Background()
	count, err := model.TaskManagements(model.TaskManagementWhere.TaskID.EQ(targetTaskId)).DeleteAllG(ctx)

	if err != nil {
		return e.WithError(err, e.InternalServerError)
	}

	if count == 0 {
		return e.WithError(err, e.DataNotFound)
	}

	return nil
}
