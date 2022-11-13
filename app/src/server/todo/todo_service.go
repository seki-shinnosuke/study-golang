package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoService struct {
}

func NewTodoService() *TodoService {
	return &TodoService{}
}

func (s *TodoService) GetTodos(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "")
}

func (s *TodoService) GetTodo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "")
}

func (s *TodoService) RegisterTodo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "")
}

func (s *TodoService) UpdateTodo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "")
}

func (s *TodoService) DeleteTodo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "")
}
