package repository

import (
	"github.com/seigi0714/go-rest-api-sample/pkg/extensions"
	"github.com/seigi0714/go-rest-api-sample/pkg/model/entity"
	"github.com/thoas/go-funk"
)

func addFields(fields []string, e entity.BaseEntity) string {
	var fieldSql string
	if len(fields) == 0 {
		for _, fd := range e.FieldsDefinition() {
			if !(fd.IsDefault) {
				continue
			}
			fieldSql = fieldSql + fd.Sql
		}
	} else {
		for _, f := range fields {
			sql := StreamOf(e.FieldsDefinition()).Filter(func(fd *entity.FieldDefinition) bool {
				return fd.Alias == f
			}).Out().([]**entity.FieldDefinition)
			fieldSql = fieldSql + sql
		}
	}

	return fieldSql
}
