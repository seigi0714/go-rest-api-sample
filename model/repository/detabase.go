package repository

import (
	"database/sql"
	"fmt"
)

var Db *sql.DB

func init() {
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"todo-app", "todo-password", "127.0.0.1:3307", "todo",
	)
	Db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
}
