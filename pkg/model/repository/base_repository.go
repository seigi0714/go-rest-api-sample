package repository

import (
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/seigi0714/go-rest-api-sample/model/entity"
)

type BaseRepository interface {
	GetTodos() (todos []entity.TodoEntity, err error)
	InsertTodo(todo entity.TodoEntity) (id int, err error)
	UpdateTodo(todo entity.TodoEntity) (err error)
	DeleteTodo(id int) (err error)
}

type baseRepository struct {
}
