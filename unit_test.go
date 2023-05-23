package sqlite

import (
	"errors"
	"log"
	"os"
	"testing"

	"github.com/cdvelop/objectdb"
)

func Test_Sqlite(t *testing.T) {
	//test.....
	root := "./test_files/"
	db_name := "test.db"

	full_path := root + db_name

	//eliminar base de datos
	if _, err := os.Stat(full_path); !errors.Is(err, os.ErrNotExist) {
		if err := os.Remove(full_path); err != nil {
			log.Printf("ERROR: %v", err)
			t.Fatal()
		}
	}

	db_sqlite := NewConnection(root, db_name, false)

	db := objectdb.Get(db_sqlite)

	db.Open()

	db.TestCrudStart(t, db_sqlite)

}
