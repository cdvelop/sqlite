package sqlite

import (
	"log"

	"github.com/cdvelop/objectdb"
)

func (d *db) TableExist(table_name string, db *objectdb.Connection) bool {

	filas, err := db.Query("SELECT name FROM sqlite_master WHERE type='table' AND name=?", table_name)
	if err != nil {
		log.Println(err)
		return false
	}
	defer filas.Close()

	return filas.Next()
}
