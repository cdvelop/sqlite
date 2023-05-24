package sqlite

import (
	"errors"
	"log"
	"os"
)

// eliminar base de datos
func (d db) DeleteDataBase() {

	full_path := d.rootFolder + d.dataBaseName

	if _, err := os.Stat(full_path); !errors.Is(err, os.ErrNotExist) {
		if err := os.Remove(full_path); err != nil {
			log.Printf("\nError al eliminar db: %v %v\n", d.dataBaseName, err)
		}
	}

}
