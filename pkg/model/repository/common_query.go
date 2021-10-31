package repository

import (
	"fmt"

	"github.com/thoas/go-funk"

	"go-rest-api-sample/pkg/extensions/slice_extensions"
	"go-rest-api-sample/pkg/model/entity"
)

func addFields(fields []string, e entity.BaseEntity) string {
	var fieldSql = ""

	fmt.Println("select field :: ", fields)
	if len(fields) == 0 {
		for _, fd := range e.FieldsDefinition() {
			if !(fd.IsDefault) {
				continue
			}
			fieldSqlJoin(&fieldSql, fd)
		}
	} else {
		sqlSlice := GetSelectedFieldsSql(e, fields)
		for _, sql := range sqlSlice {
			fieldSqlJoin(&fieldSql, sql)
		}
	}
	return fieldSql
}

func GetSelectedFieldsSql(bs entity.BaseEntity, selectedFields []string) []entity.FieldDefinition {
	return funk.Filter(bs.FieldsDefinition(), func(fd entity.FieldDefinition) bool {
		return isSelected(&fd, selectedFields)
	}).([]entity.FieldDefinition)
}

func isSelected(fd *entity.FieldDefinition, selectedFields []string) bool {
	return funk.Contains(selectedFields, fd.Alias)
}

func fieldSqlJoin(joinedFieldSql *string, joinFieldDef entity.FieldDefinition) {
	if *joinedFieldSql == "" {
		*joinedFieldSql = joinFieldDef.Sql + " as " + joinFieldDef.Alias
	} else {
		*joinedFieldSql = *joinedFieldSql + "," + joinFieldDef.Sql + " as " + joinFieldDef.Alias
	}
}

func getFieldSql(field string, e entity.BaseEntity) string {
	definition, err := slice_extensions.StreamOf(e.FieldsDefinition()).Find(func(fd *entity.FieldDefinition) bool {
		return field == fd.Alias
	})
	if err != nil {
		return ""
	}
	return definition.(entity.FieldDefinition).Sql
}
