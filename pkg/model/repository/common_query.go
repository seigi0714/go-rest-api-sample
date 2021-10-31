package repository

import (
	"github.com/thoas/go-funk"

	"go-rest-api-sample/pkg/extensions/slice_extensions"
	"go-rest-api-sample/pkg/model/entity"
)

func addFields(fields []string, e entity.BaseEntity) string {
	var fieldSql *string
	if len(fields) == 0 {
		for _, fd := range e.FieldsDefinition() {
			if !(fd.IsDefault) {
				continue
			}
			fieldSqlJoin(fieldSql, fd.Sql)
		}
	} else {
		sqlSlice := GetSelectedFieldsSql(e, fields)
		for _, sql := range sqlSlice {
			fieldSqlJoin(fieldSql, sql)
		}
	}
	return *fieldSql
}

func GetSelectedFieldsSql(bs entity.BaseEntity, selectedFields []string) []string {
	return funk.Map(
		funk.Filter(bs.FieldsDefinition(), func(fd *entity.FieldDefinition) bool {
			return isSelected(fd, selectedFields)
		}).([]entity.FieldDefinition),
		func(fd *entity.FieldDefinition) string {
			return fd.Alias
		},
	).([]string)
}

func isSelected(fd *entity.FieldDefinition, selectedFields []string) bool {
	return funk.Contains(selectedFields, fd.Alias)
}

func fieldSqlJoin(joinedFieldSql *string, joinFieldSql string) {
	if *joinedFieldSql == "" {
		*joinedFieldSql = joinFieldSql
	} else {
		*joinedFieldSql = *joinedFieldSql + " , " + joinFieldSql
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
