package sqlite

import "fmt"

func (d *db) BackupDataBase() {

	fmt.Printf("BACKUP BASE DE DATOS SQLITE NO IMPLEMENTADO\n")

}

func (d *db) RestoreDataBase(file_name string) bool {
	fmt.Println("NO IMPLEMENTADO RESTORE DB EN SQLITE")
	return false
}

func (d *db) DeleteDataBase() {

}

func (d *db) DataBaseMaintenance() {

}
