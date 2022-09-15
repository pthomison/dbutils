package postgres

import (
	"fmt"

	"github.com/pthomison/errcheck"
	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresClient struct {
	gormDB *gorm.DB

	dbName     string
	pgHost     string
	pgUser     string
	pgPort     string
	pgPassword string
}

func (c *PostgresClient) RegisterFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&c.dbName, "db-name", "", "postgres", "")
	cmd.PersistentFlags().StringVarP(&c.pgHost, "postgres-host", "", "", "")
	cmd.PersistentFlags().StringVarP(&c.pgUser, "postgres-user", "", "postgres", "")
	cmd.PersistentFlags().StringVarP(&c.pgPort, "postgres-port", "", "5432", "")
	cmd.PersistentFlags().StringVarP(&c.pgPassword, "postgres-password", "", "", "")

	cmd.MarkPersistentFlagRequired("postgres-host")
}

func (c *PostgresClient) Connect(config *gorm.Config) {

	var pw string

	if c.pgPassword == "" {
		pw = "\"\""
	} else {
		pw = c.pgPassword
	}

	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.pgHost, c.pgPort, c.pgUser, pw, c.dbName)

	db, err := gorm.Open(postgres.Open(psqlconn), config)
	errcheck.Check(err)

	c.gormDB = db
}

func (c *PostgresClient) DB() *gorm.DB {
	return c.gormDB
}
