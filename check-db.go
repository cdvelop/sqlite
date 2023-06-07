package sqlite

import (
	"log"

	"github.com/cdvelop/dbtools"
	"github.com/cdvelop/model"
	"github.com/cdvelop/objectdb"
)

// ej: root_folder: ./app_files, db_name: mydb.db
func CheckDataBase(root_folder, db_name string, mode_only_memory bool, objects ...model.Object) *objectdb.Connection {

	dba := NewConnection(root_folder, db_name, mode_only_memory)

	db := objectdb.Get(dba)
	db.Open()
	defer db.Close()

	m := dbtools.NewOperationDB(db.DB, dba)
	for _, o := range objects {

		if !dba.TableExist(o.Name, db) {
			// db.Set(dba)
			if !m.CreateOneTABLE(o) {
				log.Fatalf("Error No se logro crear tabla: %v", o.Name)
			}
		}
	}
	return db
}
