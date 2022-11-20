package todo

import (
	"database/sql"
)

type TodoUsecase struct {
	db *sql.DB
}

func NewTodoUsecase(
	db *sql.DB,
) *TodoUsecase {
	return &TodoUsecase{
		db: db,
	}
}

func (todoUsecase *TodoUsecase) GetTodos() {

}
