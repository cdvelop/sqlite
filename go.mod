module github.com/cdvelop/sqlite

go 1.20

require (
	github.com/cdvelop/model v0.0.38
	github.com/cdvelop/objectdb v0.0.38
	github.com/mattn/go-sqlite3 v1.14.17
)

require (
	github.com/cdvelop/dbtools v0.0.25 // indirect
	github.com/cdvelop/gotools v0.0.26 // indirect
	github.com/cdvelop/input v0.0.21 // indirect
	golang.org/x/text v0.11.0 // indirect
)

replace github.com/cdvelop/objectdb => ../objectdb

replace github.com/cdvelop/dbtools => ../dbtools

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/input => ../input
