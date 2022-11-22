package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	e "github.com/seki-shinnosuke/study-golang/error"
	request "github.com/seki-shinnosuke/study-golang/model/request/todo"
	response "github.com/seki-shinnosuke/study-golang/model/response/todo"
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
	todos, err := ctrl.todoUsecase.GetTodos()
	if err != nil {
		appError := e.Cast(err)
		ctx.JSON(appError.StatusCode, appError)
		return
	}

	ctx.JSON(http.StatusOK, todos)
}

func (ctrl *TodoController) GetTodo(ctx *gin.Context) {
	var uriParam request.UriParam
	if err := ctx.ShouldBindUri(&uriParam); err != nil {
		ctx.JSON(e.InvalidRequestParameters.StatusCode, e.InvalidRequestParameters)
		return
	}

	todo, err := ctrl.todoUsecase.GetTodo(uriParam.Id)
	if err != nil {
		appError := e.Cast(err)
		ctx.JSON(appError.StatusCode, appError)
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

func (ctrl *TodoController) RegisterTodo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "")
}

func (ctrl *TodoController) UpdateTodo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "")
}

func (ctrl *TodoController) DeleteTodo(ctx *gin.Context) {
	var uriParam request.UriParam
	if err := ctx.ShouldBindUri(&uriParam); err != nil {
		ctx.JSON(e.InvalidRequestParameters.StatusCode, e.InvalidRequestParameters)
		return
	}

	err := ctrl.todoUsecase.DeleteTodo(uriParam.Id)
	if err != nil {
		appError := e.Cast(err)
		ctx.JSON(appError.StatusCode, appError)
		return
	}

	ctx.JSON(http.StatusOK, response.TaskDeleteResponse{})
}
