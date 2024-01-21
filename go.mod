module github.com/cdvelop/sqlite

go 1.20

require (
	github.com/cdvelop/model v0.0.119
	github.com/cdvelop/objectdb v0.0.116
	github.com/mattn/go-sqlite3 v1.14.19
)

require (
	github.com/cdvelop/dbtools v0.0.82 // indirect
	github.com/cdvelop/input v0.0.83 // indirect
	github.com/cdvelop/maps v0.0.8 // indirect
	github.com/cdvelop/strings v0.0.9 // indirect
	github.com/cdvelop/timeserver v0.0.33 // indirect
	github.com/cdvelop/timetools v0.0.37 // indirect
	github.com/cdvelop/unixid v0.0.49 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/objectdb => ../objectdb

replace github.com/cdvelop/dbtools => ../dbtools
