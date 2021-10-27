package entity

type BaseEntity interface {
	Teble() string
	PrimaryKey() string
	FieldsDefinition() []FieldDefinition
}

type FieldDefinition struct {
	Alias     string
	IsDefault bool
	Sql       string
}
