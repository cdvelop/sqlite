package sqlite

import "fmt"

// SQLParameters string  //ej: postgres: "$" sqlite "?",
func (d *db) PlaceHolders(index ...uint8) string {
	return `?`
}

//SetListSyntax "%v=?", k
func (d *db) SetListSyntax(key string, i byte, set *[]string) {
	*set = append(*set, fmt.Sprintf("%v=?", key))
}

func (d *db) TotalValuesSyntax(fields map[string]string) string {
	return "?"
}

func (d *db) MakeSqInsertSyntax(i *byte, setValue *[]string) {
	*setValue = append(*setValue, "?")
}
