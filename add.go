package sqlite

import (
	"fmt"

	"github.com/cdvelop/dbtools"
	"github.com/cdvelop/model"
	"github.com/cdvelop/objectdb"
	_ "github.com/mattn/go-sqlite3"
)

// root_folder ej: "./app_files" db_name ej: mi_proyecto.db min 4 caracteres
func NewConnection(root_folder, db_name string, mode_only_memory bool, tables ...model.Object) *objectdb.Connection {

	if db_name == "" || len(db_name) < 4 {
		showErrorAndExit("NOMBRE BASE DE DATOS INCORRECTO")
	}

	dba := db{
		rootFolder:   root_folder + "/",
		dataBaseName: db_name,
	}

	db := objectdb.Get(&dba)

	// chequear tablas base de datos
	for _, t := range tables {
		if !dba.TableExist(t.Name, db) {
			if !dbtools.CreateOneTABLE(db, t) {
				showErrorAndExit(fmt.Sprintf("no se logro crear tabla: %v", t.Name))
			}
		}
	}
	return db
}
