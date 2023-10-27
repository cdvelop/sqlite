module github.com/cdvelop/sqlite

go 1.20

require (
	github.com/cdvelop/model v0.0.61
	github.com/cdvelop/objectdb v0.0.73
	github.com/mattn/go-sqlite3 v1.14.17
)

require (
	github.com/cdvelop/dbtools v0.0.50 // indirect
	github.com/cdvelop/gotools v0.0.48 // indirect
	github.com/cdvelop/input v0.0.43 // indirect
	github.com/cdvelop/timeserver v0.0.8 // indirect
	github.com/cdvelop/timetools v0.0.9 // indirect
	github.com/cdvelop/unixid v0.0.9 // indirect
	golang.org/x/text v0.13.0 // indirect
)

replace github.com/cdvelop/objectdb => ../objectdb

replace github.com/cdvelop/unixid => ../unixid

replace github.com/cdvelop/timeserver => ../timeserver

replace github.com/cdvelop/timetools => ../timetools

replace github.com/cdvelop/dbtools => ../dbtools

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/input => ../input
