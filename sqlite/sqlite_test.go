package sqlite

import (
	"os"
	"testing"

	"github.com/pthomison/dbutils"
	"github.com/pthomison/errcheck"
)

func TestSQLiteConnectAndWriteDB(t *testing.T) {
	dbLocation := "./gorm.db"

	client := &SQLiteClient{
		sqliteFile: dbLocation,
	}

	dbutils.ConnectAndWriteDBTest(t, client)

	err := os.Remove(dbLocation)
	errcheck.CheckTest(err, t)
}
