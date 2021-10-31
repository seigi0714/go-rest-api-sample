package entity

type BaseEntity interface {
	Table() string
	PrimaryKey() string
	FieldsDefinition() []FieldDefinition
}

type FieldDefinition struct {
	Alias     string
	IsDefault bool
	Sql       string
}
