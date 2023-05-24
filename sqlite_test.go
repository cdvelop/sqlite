package sqlite_test

import (
	"testing"

	"github.com/cdvelop/objectdb"
	"github.com/cdvelop/sqlite"
)

func Test_Sqlite(t *testing.T) {
	//test.....
	root := "./test_files"
	db_name := "test.db"

	db_sqlite := sqlite.NewConnection(root, db_name, false)

	db_sqlite.DeleteDataBase()

	db := objectdb.Get(db_sqlite)

	db.Open()

	db.TestCrudStart(t, db_sqlite)

}
