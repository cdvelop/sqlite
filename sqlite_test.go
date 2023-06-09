package sqlite_test

import (
	"testing"

	"github.com/cdvelop/sqlite"
)

func Test_Sqlite(t *testing.T) {
	//test.....
	root := "./test_files"
	db_name := "test.db"

	db := sqlite.NewConnection(root, db_name, false)

	db.TestCrudStart(t)

}
