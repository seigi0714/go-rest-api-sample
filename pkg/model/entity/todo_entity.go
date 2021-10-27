package entity

type TodoEntity struct {
	Id      int
	Title   string
	Content string
}

func (te *TodoEntity) Table() string {
	return "todo"
}

func (te *TodoEntity) PrimaryKey() string {
	return "id"
}

func (te *TodoEntity) FieldsDefinition() []FieldDefinition {
	return []FieldDefinition{
		{"id", true, "todo.id"},
		{"title", true, "todo.title"},
		{"content", true, "todo.content"},
	}

}
