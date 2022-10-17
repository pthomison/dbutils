package dbutils

import (
	"fmt"

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

func Migrate(c DBClient, obj interface{}) {
	err := c.DB().AutoMigrate(obj)
	errcheck.Check(err)
}

func Create[T any](c DBClient, objs []T) {
	fmt.Println(objs)
	result := c.DB().Create(&objs)
	errcheck.Check(result.Error)
	fmt.Printf("%+v\n", result)

}

func CreateOrOverwrite[T any](c DBClient, objs []T) {
	result := c.DB().Clauses(clause.OnConflict{UpdateAll: true}).Create(objs)
	errcheck.Check(result.Error)
}

func DeleteAll(c DBClient, obj interface{}) {
	result := c.DB().Where("1 = 1").Unscoped().Delete(obj)
	errcheck.Check(result.Error)
}
