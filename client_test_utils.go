package dbutils

import (
	"errors"
	"fmt"
	"testing"

	"github.com/pthomison/errcheck"
	"gorm.io/gorm"
)

type TestData struct {
	gorm.Model

	IntData    int
	StringData string
	BoolData   bool
}

func (td *TestData) Compare(comp *TestData) bool {
	intComp := td.IntData == comp.IntData
	stringComp := td.StringData == comp.StringData
	boolComp := td.BoolData == comp.BoolData

	return intComp && stringComp && boolComp
}

func ConnectAndWriteDBTest(t *testing.T, dbc DBClient) {
	var data []TestData
	var fetchedData []TestData

	dbc.Connect(&gorm.Config{})

	Migrate(dbc, &TestData{})

	data = []TestData{
		TestData{
			IntData:    10,
			StringData: "10",
			BoolData:   true,
		},
	}

	Create(dbc, data)

	fetchedData = Query[TestData](dbc.DB())

	fmt.Println(fetchedData)

	if len(fetchedData) == 0 || !data[0].Compare(&fetchedData[0]) {
		errcheck.CheckTest(errors.New("injected data doesn't match retrieved data"), t)
	}

	DeleteAll(dbc, &TestData{})

	fetchedData = Query[TestData](dbc.DB())

	if len(fetchedData) != 0 {
		errcheck.CheckTest(errors.New("delete failed"), t)
	}

}
