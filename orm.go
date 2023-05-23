package sqlite

func (d *db) DataBasEngine() string {
	return "sqlite3"
}

// ej "mydb"
func (d db) DataBaseName() string {
	return d.dataBaseName
}

// ConnectionString formato cadena de conexión
func (d *db) ConnectionString() string {
	return d.rootFolder + d.dataBaseName
}

func (d *db) SQLTableInfo() string {
	return "SELECT name FROM PRAGMA_TABLE_INFO('%v');"
}

func (d *db) SQLColumName() string {
	return "name"
}

func (d *db) SQLDropTable() string {
	return "DROP TABLE IF EXISTS %v;" //sql de eliminación de tabla
}
