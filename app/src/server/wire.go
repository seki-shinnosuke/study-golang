//go:build wireinject
// +build wireinject

package server

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/seki-shinnosuke/study-golang/config"
	todoctrl "github.com/seki-shinnosuke/study-golang/controller/todo"
	todouc "github.com/seki-shinnosuke/study-golang/usecase/todo"
)

func InitializeService(
	configAPIServer *config.APIServer,
	db *sql.DB,
) *Server {
	wire.Build(
		todoctrl.NewTodoController,
		todouc.NewTodoUsecase,
		NewServer,
	)
	return &Server{}
}
