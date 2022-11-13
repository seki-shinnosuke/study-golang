//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"
	"github.com/seki-shinnosuke/study-golang/config"
	"github.com/seki-shinnosuke/study-golang/server/todo"
)

func InitializeService(configAPIServer *config.APIServer) *Server {
	wire.Build(
		todo.NewTodoService,
		NewServer,
	)
	return &Server{}
}
