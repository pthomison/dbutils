package sqlite

import (
	"github.com/glebarez/sqlite" // pure go driver
	"github.com/pthomison/errcheck"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

type SQLiteClient struct {
	gormDB *gorm.DB

	SQLiteFile string
}

func New(sqliteFile string) *SQLiteClient {
	client := &SQLiteClient{
		SQLiteFile: sqliteFile,
	}

	client.Connect(&gorm.Config{})

	return client
}

func (c *SQLiteClient) RegisterFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&c.SQLiteFile, "sqlite-file", "", "gorm.db", "")
	cmd.MarkPersistentFlagRequired("sqlite-file")
}

func (c *SQLiteClient) Connect(config *gorm.Config) {
	db, err := gorm.Open(sqlite.Open(c.SQLiteFile), config)
	errcheck.Check(err)

	c.gormDB = db
}

func (c *SQLiteClient) DB() *gorm.DB {
	return c.gormDB
}
