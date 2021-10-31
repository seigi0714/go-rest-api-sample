package repository

import (
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"go-rest-api-sample/pkg/model/entity"
)

type TodoRepository interface {
	GetTodo(id int, fields []string) (result map[string]interface{}, err error)
	GetTodos(fields []string) (todos []map[string]interface{}, err error)
	InsertTodo(todo entity.TodoEntity) (id int, err error)
	UpdateTodo(todo entity.TodoEntity) (err error)
	DeleteTodo(id int) (err error)
}

type todoRepository struct {
}

func NewTodoRepository() TodoRepository {
	return &todoRepository{}
}

func (tr *todoRepository) GetTodo(id int, fields []string) (result map[string]interface{}, err error) {
	// TODO:: 共通のSQL文なので関数化
	e := entity.NewTodoEntity()
	fieldSql := "SELECT " + addFields(fields, e)
	from := " FROM todo"
	where := " WHERE id = " + strconv.Itoa(id)

	sql := fieldSql + from + where
	rows, _ := Db.
		Query(sql)
	cols, err := rows.Columns()
	for rows.Next() {
		var row = make([]interface{}, len(cols))
		var rowp = make([]interface{}, len(cols))
		for i := 0; i < len(cols); i++ {
			rowp[i] = &row[i]
		}

		rows.Scan(rowp...)
		rowMap := make(map[string]interface{})
		for i, col := range cols {
			switch row[i].(type) {
			case []byte:
				row[i] = string(row[i].([]byte))
				num, err := strconv.Atoi(row[i].(string))
				if err == nil {
					row[i] = num
				}
			}
			rowMap[col] = row[i]
		}
		result = rowMap
	}
	return
}

func (tr *todoRepository) GetTodos(fields []string) (result []map[string]interface{}, err error) {
	// TODO:: 共通のSQL文なので関数化
	e := entity.NewTodoEntity()
	fieldSql := "SELECT " + addFields(fields, e)
	from := " FROM todo"
	sort := " ORDER BY id"

	sql := fieldSql + from + sort
	rows, _ := Db.
		Query(sql)
	cols, err := rows.Columns()
	for rows.Next() {
		var row = make([]interface{}, len(cols))
		var rowp = make([]interface{}, len(cols))
		for i := 0; i < len(cols); i++ {
			rowp[i] = &row[i]
		}

		rows.Scan(rowp...)

		rowMap := make(map[string]interface{})
		for i, col := range cols {
			switch row[i].(type) {
			case []byte:
				row[i] = string(row[i].([]byte))
				num, err := strconv.Atoi(row[i].(string))
				if err == nil {
					row[i] = num
				}
			}
			rowMap[col] = row[i]
		}
		result = append(result, rowMap)
	}
	return
}

func (tr *todoRepository) InsertTodo(todo entity.TodoEntity) (id int, err error) {
	_, err = Db.Exec("INSERT INTO todo (title, content) VALUES (?, ?)", todo.Title, todo.Content)
	if err != nil {
		log.Print(err)
		return
	}
	err = Db.QueryRow("SELECT id FROM todo ORDER BY id DESC LIMIT 1").Scan(&id)
	return
}

func (tr *todoRepository) UpdateTodo(todo entity.TodoEntity) (err error) {
	_, err = Db.Exec("UPDATE todo SET title = ?, content = ? WHERE id = ?", todo.Title, todo.Content, todo.Id)
	return
}

func (tr *todoRepository) DeleteTodo(id int) (err error) {
	_, err = Db.Exec("DELETE FROM todo WHERE id = ?", id)
	return
}
