package server

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/seki-shinnosuke/study-golang/config"
	"github.com/seki-shinnosuke/study-golang/controller/todo"
	e "github.com/seki-shinnosuke/study-golang/error"
	"github.com/seki-shinnosuke/study-golang/util/logger"
)

type Server struct {
	configAPIServer *config.APIServer
	Gin             *gin.Engine
	todoController  *todo.TodoController
}

func NewServer(
	configAPIServer *config.APIServer,
	todoController *todo.TodoController,
) *Server {
	gin.SetMode(configAPIServer.GinMode)
	server := &Server{
		configAPIServer: configAPIServer,
		Gin:             gin.New(),
		todoController:  todoController,
	}
	server.setCors()
	server.setRouting()
	return server
}

func (server *Server) setCors() {
	corsOrigins := []string{}
	if len(server.configAPIServer.CorsOrigins) != 0 {
		corsOrigins = strings.Split(server.configAPIServer.CorsOrigins, ",")
	}
	server.Gin.Use(
		cors.New(cors.Config{
			AllowOrigins:     corsOrigins,
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Access-Control-Allow-Headers", "Content-Length", "Content-Type", "Authorization"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}),
		gin.LoggerWithConfig(logger.CustomGinLogger([]string{"/health"})),
		gin.Recovery(),
	)
}

func (server *Server) setRouting() {
	server.Gin.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "OK")
	})
	apiV1 := server.Gin.Group("/api/v1")
	{
		apiV1.GET("/tasks", server.todoController.GetTodos)
		apiV1.GET("/tasks/:id", server.todoController.GetTodo)
		apiV1.POST("/tasks", server.todoController.RegisterTodo)
		apiV1.PUT("/tasks/:id", server.todoController.UpdateTodo)
		apiV1.DELETE("/tasks/:id", server.todoController.DeleteTodo)
	}
	server.Gin.NoRoute(func(c *gin.Context) {
		c.JSON(e.NotFound.StatusCode, e.NotFound)
	})
}

func (server *Server) Run() error {
	return server.Gin.Run(":" + server.configAPIServer.AppApiPort)
}
