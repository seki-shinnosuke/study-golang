package controller

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/seki-shinnosuke/study-golang/config"
)

type Routing struct {
	configAPIServer *config.APIServer
	Gin             *gin.Engine
}

func NewRouting(
	configAPIServer *config.APIServer,
) *Routing {
	r := &Routing{
		configAPIServer: configAPIServer,
		Gin:             gin.Default(),
	}
	r.setCors()
	r.setRouting()
	return r
}

func (r *Routing) setCors() {
	r.Gin.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:30000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}

func (r *Routing) setRouting() {
	r.Gin.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "OK")
	})
	// apiV1 := r.Gin.Group("/api/v1")
	// {

	// }
	r.Gin.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "The page not found"})
	})
}

func (r *Routing) Run() error {
	return r.Gin.Run(":" + r.configAPIServer.AppApiPort)
}
