package dbutils

import (
	"github.com/pthomison/errcheck"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DBClient interface {
	Connect(*gorm.Config)
	DB() *gorm.DB
	RegisterFlags(cmd *cobra.Command)
}

func SelectAll[T any](c DBClient, columns []string) []T {
	var output []T
	result := c.DB().Select(columns).Find(&output)
	errcheck.Check(result.Error)

	return output
}

func SelectWhere[T any](c DBClient, columns []string, whereQuery interface{}, whereArgs ...interface{}) []T {
	var output []T
	result := c.DB().Where(whereQuery, whereArgs).Select(columns).Find(&output)
	errcheck.Check(result.Error)

	return output
}

func Create[T any](c DBClient, objs []T) {
	result := c.DB().Create(objs)
	errcheck.Check(result.Error)
}

func CreateOrOverwrite[T any](c DBClient, objs []T) {
	result := c.DB().Clauses(clause.OnConflict{UpdateAll: true}).Create(objs)
	errcheck.Check(result.Error)
}
