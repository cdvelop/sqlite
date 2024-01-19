package sqlite_test

import (
	"testing"

	"github.com/cdvelop/sqlite"
)

func Test_SqlitePersistenMode(t *testing.T) {
	//test.....
	root := "./test_files"
	db_name := "test.db"

	db := sqlite.NewConnection(root, db_name, false)

	db.TestCrudStart(t)

}

func Test_SqliteMemoryMode(t *testing.T) {
	db := sqlite.NewConnection("", "test_memory", true)
	db.TestCrudStart(t)

}
