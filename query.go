package dbutils

import (
	"github.com/pthomison/errcheck"
	"gorm.io/gorm"
)

const (
	DEFAULT_WHERE = "TRUE"
	DEFAULT_ORDER = "ID ASC"
)

func Query[T any](query *gorm.DB) []T {
	results := []T{}

	ret := query.Find(&results)

	errcheck.Check(ret.Error)

	return results
}

func SelectAll[T any](c DBClient) []T {
	return Query[T](c.DB())
}
