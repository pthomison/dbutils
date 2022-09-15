package sqlite

import (
	"github.com/pthomison/errcheck"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLiteClient struct {
	gormDB *gorm.DB

	sqliteFile string
}

func (c *SQLiteClient) RegisterFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&c.sqliteFile, "sqlite-file", "", "gorm.db", "")
	cmd.MarkPersistentFlagRequired("sqlite-file")
}

func (c *SQLiteClient) Connect(config *gorm.Config) {
	db, err := gorm.Open(sqlite.Open(c.sqliteFile), config)
	errcheck.Check(err)

	c.gormDB = db
}

func (c *SQLiteClient) DB() *gorm.DB {
	return c.gormDB
}
