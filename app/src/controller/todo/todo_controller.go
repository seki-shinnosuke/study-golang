package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	todouc "github.com/seki-shinnosuke/study-golang/usecase/todo"
)

type TodoController struct {
	todoUsecase *todouc.TodoUsecase
}

func NewTodoController(
	todoUsecase *todouc.TodoUsecase,
) *TodoController {
	return &TodoController{
		todoUsecase: todoUsecase,
	}
}

func (s *TodoController) GetTodos(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "")
}

func (s *TodoController) GetTodo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "")
}

func (s *TodoController) RegisterTodo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "")
}

func (s *TodoController) UpdateTodo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "")
}

func (s *TodoController) DeleteTodo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "")
}
