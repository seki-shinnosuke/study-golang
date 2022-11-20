//go:build wireinject
// +build wireinject

package server

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/seki-shinnosuke/study-golang/config"
	"github.com/seki-shinnosuke/study-golang/server/todo"
)

func InitializeService(
	configAPIServer *config.APIServer,
	db *sql.DB,
) *Server {
	wire.Build(
		todo.NewTodoService,
		NewServer,
	)
	return &Server{}
}
