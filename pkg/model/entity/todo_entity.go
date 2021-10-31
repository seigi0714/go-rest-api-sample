package entity

type TodoEntity struct {
	Id      int    `db:"todo.id"`
	Title   string `db:"todo.title"`
	Content string `db:"todo.content"`
}

type todoEntity struct {
}

func NewTodoEntity() *todoEntity {
	return &todoEntity{}
}

func (te *todoEntity) Table() string {
	return "todo"
}

func (te *todoEntity) PrimaryKey() string {
	return "id"
}

func (te *todoEntity) FieldsDefinition() []FieldDefinition {
	return []FieldDefinition{
		{"id", true, "todo.id"},
		{"title", true, "todo.title"},
		{"content", true, "todo.content"},
	}
}
