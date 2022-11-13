package server

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/seki-shinnosuke/study-golang/config"
	"github.com/seki-shinnosuke/study-golang/server/todo"
)

type Server struct {
	configAPIServer *config.APIServer
	Gin             *gin.Engine
	todoService     *todo.TodoService
}

func NewServer(
	configAPIServer *config.APIServer,
	todoService *todo.TodoService,
) *Server {
	gin.SetMode(configAPIServer.GinMode)
	server := &Server{
		configAPIServer: configAPIServer,
		Gin:             gin.New(),
		todoService:     todoService,
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
		gin.LoggerWithConfig(gin.LoggerConfig{
			SkipPaths: []string{"/health"},
		}))
}

func (server *Server) setRouting() {
	server.Gin.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "OK")
	})
	apiV1 := server.Gin.Group("/api/v1")
	{
		apiV1.GET("/tasks", server.todoService.GetTodos)
		apiV1.GET("/tasks:id", server.todoService.GetTodo)
		apiV1.POST("/tasks", server.todoService.RegisterTodo)
		apiV1.PUT("/tasks:id", server.todoService.UpdateTodo)
		apiV1.DELETE("/tasks:id", server.todoService.DeleteTodo)
	}
	server.Gin.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "The page not found"})
	})
}

func (server *Server) Run() error {
	return server.Gin.Run(":" + server.configAPIServer.AppApiPort)
}
