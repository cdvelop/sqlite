package sqlite

import (
	"github.com/cdvelop/model"
	"github.com/cdvelop/objectdb"
	_ "github.com/mattn/go-sqlite3"
)

// root_folder ej: "./app_files" db_name ej: mi_proyecto.db min 4 caracteres
func NewConnection(root_folder, db_name string, mode_only_memory bool, tables ...*model.Object) *objectdb.Connection {

	if db_name == "" || len(db_name) < 4 {
		showErrorAndExit("NOMBRE BASE DE DATOS INCORRECTO")
	}

	dba := db{
		rootFolder:   root_folder + "/",
		dataBaseName: db_name,
	}

	db := objectdb.Get(&dba)

	// chequear tablas base de datos
	db.CreateTablesInDB(tables, func(err string) {
		if err != "" {
			showErrorAndExit(err)
		}
	})

	return db
}
