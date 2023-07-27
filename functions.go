package sqlite

import "fmt"

// SQLParameters string  //ej: postgres: "$" sqlite "?",
func (db) PlaceHolders(index ...uint8) string {
	return `?`
}

//SetListSyntax "%v=?", k
func (db) SetListSyntax(key string, i byte, set *[]string) {
	*set = append(*set, fmt.Sprintf("%v=?", key))
}

func (db) TotalValuesSyntax(fields map[string]string) string {
	return "?"
}

func (db) MakeSqInsertSyntax(i *byte, setValue *[]string) {
	*setValue = append(*setValue, "?")
}

func (db) DropTable() string {
	return "DROP TABLE IF EXISTS %v"
}

func (db) SQLTableExist() string {
	return "SELECT name FROM sqlite_master WHERE type='table' AND name=?"
}
