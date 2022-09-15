package postgres

import (
	"testing"

	"github.com/pthomison/dbutils"
)

func TestSQLiteConnectAndWriteDB(t *testing.T) {
	client := &PostgresClient{
		dbName:     "postgres",
		pgHost:     "127.0.0.1",
		pgUser:     "pthomison",
		pgPort:     "5432",
		pgPassword: "",
	}

	dbutils.ConnectAndWriteDBTest(t, client)
}

// type PostgresClient struct {
// 	gormDB *gorm.DB

// 	dbName     string
// 	pgHost     string
// 	pgUser     string
// 	pgPort     string
// 	pgPassword string
// }
