package server

import (
	"net/http"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/seki-shinnosuke/study-golang/config"
	"github.com/seki-shinnosuke/study-golang/controller/todo"
	e "github.com/seki-shinnosuke/study-golang/error"
	gq "github.com/seki-shinnosuke/study-golang/server/graphql"
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
			AllowHeaders:     []string{"Access-Control-Allow-Headers", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}),
		gin.LoggerWithConfig(logger.CustomGinLogger([]string{"/health"})),
		gin.Recovery(),
	)
}

func graphqlHandler() gin.HandlerFunc {
	h := handler.New(gq.NewExecutableSchema(gq.Config{Resolvers: &gq.Resolver{}}))
	h.AddTransport(transport.POST{})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func (server *Server) setRouting() {
	server.Gin.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "OK")
	})
	// GraphQL
	graphql := server.Gin.Group("/graphql")
	{
		graphql.POST("", graphqlHandler())
		graphql.GET("/playground", playgroundHandler())
	}
	//REST API
	apiV1 := server.Gin.Group("/api/v1")
	{
		apiV1.GET("/tasks", server.todoController.GetTasks)
		apiV1.GET("/tasks/:taskId", server.todoController.GetTask)
		apiV1.POST("/tasks", server.todoController.RegisterTask)
		apiV1.PUT("/tasks/:taskId", server.todoController.UpdateTask)
		apiV1.DELETE("/tasks/:taskId", server.todoController.DeleteTask)
	}
	server.Gin.NoRoute(func(c *gin.Context) {
		c.JSON(e.NotFound.StatusCode, e.NotFound)
	})
}

func (server *Server) Run() error {
	return server.Gin.Run(":" + server.configAPIServer.AppApiPort)
}
