package sqlite

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// root_folder ej: "./app_files" database_name ej: mi_proyecto.db min 4 caracteres
func NewConnection(root_folder, database_name string, mode_only_memory bool) *db {

	if database_name == "" || len(database_name) < 4 {
		log.Fatalln("NOMBRE BASE DE DATOS INCORRECTO")
	}

	d := db{
		rootFolder:   root_folder + "/",
		dataBaseName: database_name,
	}

	return &d
}
