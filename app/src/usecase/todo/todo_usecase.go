package todo

import (
	"context"

	models "github.com/seki-shinnosuke/study-golang/model/db"
	"github.com/seki-shinnosuke/study-golang/util/logger"
)

type TodoUsecase struct {
}

func NewTodoUsecase() *TodoUsecase {
	return &TodoUsecase{}
}

func (uc *TodoUsecase) GetTodos() {
	ctx := context.Background()
	todos, err := models.TaskManagements().AllG(ctx)

	if err != nil {

	}
	logger.Info("%s", todos)
}
