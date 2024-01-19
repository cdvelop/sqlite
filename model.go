package sqlite

// db formato cadena conexión
type db struct {
	rootFolder   string // carpeta donde se alojara la db ej: "./mi_db"
	dataBaseName string //nombre de la base de datos ej: mi_proyecto.db

	modeMemory bool //solo en memoria
}
