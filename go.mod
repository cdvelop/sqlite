module github.com/cdvelop/sqlite

go 1.20

require (
	github.com/cdvelop/model v0.0.111
	github.com/cdvelop/objectdb v0.0.114
	github.com/mattn/go-sqlite3 v1.14.19
)

require (
	github.com/cdvelop/dbtools v0.0.81 // indirect
	github.com/cdvelop/input v0.0.83 // indirect
	github.com/cdvelop/maps v0.0.8 // indirect
	github.com/cdvelop/strings v0.0.9 // indirect
	github.com/cdvelop/timeserver v0.0.32 // indirect
	github.com/cdvelop/timetools v0.0.34 // indirect
	github.com/cdvelop/unixid v0.0.49 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/objectdb => ../objectdb
