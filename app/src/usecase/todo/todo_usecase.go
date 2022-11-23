package todo

import (
	"context"
	"database/sql"

	e "github.com/seki-shinnosuke/study-golang/error"
	model "github.com/seki-shinnosuke/study-golang/model/db"
	request "github.com/seki-shinnosuke/study-golang/model/request/todo"
	response "github.com/seki-shinnosuke/study-golang/model/response/todo"
	"github.com/seki-shinnosuke/study-golang/util/logger"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type TodoUsecase struct {
}

func NewTodoUsecase() *TodoUsecase {
	return &TodoUsecase{}
}

func (uc *TodoUsecase) GetTasks() (*response.TodoResponse, error) {
	ctx := context.Background()
	dbTasks, err := model.TaskManagements().AllG(ctx)

	if err != nil {
		return nil, e.WithError(err, e.InternalServerError)
	}

	tasks := make([]response.Task, 0, len(dbTasks))
	for _, dbTask := range dbTasks {
		tasks = append(tasks, response.Task{
			TaskId:       dbTask.TaskID,
			PersonName:   dbTask.PersonName,
			TaskName:     dbTask.TaskName,
			DeadlineDate: dbTask.DeadlineDate,
			TaskStatus:   dbTask.TaskStatus,
		})
	}

	return &response.TodoResponse{Tasks: tasks}, nil
}

func (uc *TodoUsecase) GetTask(targetTaskId int) (*response.TodoDetailResponse, error) {
	ctx := context.Background()
	dbTask, err := model.FindTaskManagementG(ctx, targetTaskId)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, e.WithError(err, e.DataNotFound)
		}
		return nil, e.WithError(err, e.InternalServerError)
	}

	return &response.TodoDetailResponse{Task: response.Task{
		TaskId:       dbTask.TaskID,
		PersonName:   dbTask.PersonName,
		TaskName:     dbTask.TaskName,
		DeadlineDate: dbTask.DeadlineDate,
		TaskStatus:   dbTask.TaskStatus,
	}}, nil
}

func (uc *TodoUsecase) RegisterTask(task request.Task) (*response.TodoDetailResponse, error) {
	registerTask := model.TaskManagement{
		PersonName:   task.PersonName,
		TaskName:     task.TaskName,
		DeadlineDate: task.DeadlineDate,
		TaskStatus:   task.TaskStatus,
	}

	ctx := context.Background()
	err := registerTask.InsertG(ctx, boil.Infer())

	if err != nil {
		logger.Error(" %v", err)
		return nil, e.WithError(err, e.InternalServerError)
	}

	return &response.TodoDetailResponse{Task: response.Task{
		TaskId:       registerTask.TaskID,
		PersonName:   registerTask.PersonName,
		TaskName:     registerTask.TaskName,
		DeadlineDate: registerTask.DeadlineDate,
		TaskStatus:   registerTask.TaskStatus,
	}}, nil
}

func (uc *TodoUsecase) UpdateTask(targetTaskId int, task request.Task) (*response.TodoDetailResponse, error) {
	ctx := context.Background()
	dbUpdateTask, err := model.FindTaskManagementG(ctx, targetTaskId)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, e.WithError(err, e.DataNotFound)
		}
		return nil, e.WithError(err, e.InternalServerError)
	}

	dbUpdateTask.PersonName = task.PersonName
	dbUpdateTask.TaskName = task.TaskName
	dbUpdateTask.DeadlineDate = task.DeadlineDate
	dbUpdateTask.TaskStatus = task.TaskStatus

	count, err := dbUpdateTask.UpdateG(ctx, boil.Infer())

	if err != nil || count != 1 {
		logger.Error(" %v", err)
		return nil, e.WithError(err, e.InternalServerError)
	}

	return &response.TodoDetailResponse{Task: response.Task{
		TaskId:       dbUpdateTask.TaskID,
		PersonName:   dbUpdateTask.PersonName,
		TaskName:     dbUpdateTask.TaskName,
		DeadlineDate: dbUpdateTask.DeadlineDate,
		TaskStatus:   dbUpdateTask.TaskStatus,
	}}, nil
}

func (uc *TodoUsecase) DeleteTask(targetTaskId int) error {
	ctx := context.Background()
	count, err := model.TaskManagements(model.TaskManagementWhere.TaskID.EQ(targetTaskId)).DeleteAllG(ctx)

	if err != nil {
		logger.Error(" %v", err)
		return e.WithError(err, e.InternalServerError)
	}

	if count == 0 {
		return e.WithError(err, e.DataNotFound)
	}

	return nil
}
