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

func (ctrl *TodoController) GetTasks(ctx *gin.Context) {
	response, err := ctrl.todoUsecase.GetTasks()

	if err != nil {
		appError := e.Cast(err)
		ctx.JSON(appError.StatusCode, appError)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (ctrl *TodoController) GetTask(ctx *gin.Context) {
	var uriParam request.UriParam

	if err := ctx.ShouldBindUri(&uriParam); err != nil {
		ctx.JSON(e.InvalidRequestParameters.StatusCode, e.InvalidRequestParameters)
		return
	}

	response, err := ctrl.todoUsecase.GetTask(uriParam.TaskId)

	if err != nil {
		appError := e.Cast(err)
		ctx.JSON(appError.StatusCode, appError)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (ctrl *TodoController) RegisterTask(ctx *gin.Context) {
	var requestBody request.Task

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(e.InvalidRequestParameters.StatusCode, e.InvalidRequestParameters)
		return
	}

	response, err := ctrl.todoUsecase.RegisterTask(requestBody)

	if err != nil {
		appError := e.Cast(err)
		ctx.JSON(appError.StatusCode, appError)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (ctrl *TodoController) UpdateTask(ctx *gin.Context) {
	var uriParam request.UriParam

	if err := ctx.ShouldBindUri(&uriParam); err != nil {
		ctx.JSON(e.InvalidRequestParameters.StatusCode, e.InvalidRequestParameters)
		return
	}

	var requestBody request.Task

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(e.InvalidRequestParameters.StatusCode, e.InvalidRequestParameters)
		return
	}

	response, err := ctrl.todoUsecase.UpdateTask(uriParam.TaskId, requestBody)

	if err != nil {
		appError := e.Cast(err)
		ctx.JSON(appError.StatusCode, appError)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (ctrl *TodoController) DeleteTask(ctx *gin.Context) {
	var uriParam request.UriParam
	if err := ctx.ShouldBindUri(&uriParam); err != nil {
		ctx.JSON(e.InvalidRequestParameters.StatusCode, e.InvalidRequestParameters)
		return
	}

	err := ctrl.todoUsecase.DeleteTask(uriParam.TaskId)
	if err != nil {
		appError := e.Cast(err)
		ctx.JSON(appError.StatusCode, appError)
		return
	}

	ctx.JSON(http.StatusOK, response.TodoDeleteResponse{})
}
