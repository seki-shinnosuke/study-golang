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

func (ctrl *TodoController) GetTodos(ctx *gin.Context) {
	todos, _ := ctrl.todoUsecase.GetTodos()
	ctx.JSON(http.StatusOK, todos)
}

func (ctrl *TodoController) GetTodo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "")
}

func (ctrl *TodoController) RegisterTodo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "")
}

func (ctrl *TodoController) UpdateTodo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "")
}

func (ctrl *TodoController) DeleteTodo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "")
}
