module github.com/cdvelop/sqlite

go 1.20

require (
	github.com/cdvelop/model v0.0.55
	github.com/cdvelop/objectdb v0.0.67
	github.com/mattn/go-sqlite3 v1.14.17
)

require (
	github.com/cdvelop/dbtools v0.0.44 // indirect
	github.com/cdvelop/gotools v0.0.43 // indirect
	github.com/cdvelop/input v0.0.37 // indirect
	github.com/cdvelop/timeserver v0.0.3 // indirect
	github.com/cdvelop/timetools v0.0.4 // indirect
	github.com/cdvelop/unixid v0.0.5 // indirect
	golang.org/x/text v0.13.0 // indirect
)

replace github.com/cdvelop/objectdb => ../objectdb

replace github.com/cdvelop/timeserver => ../timeserver

replace github.com/cdvelop/timetools => ../timetools

replace github.com/cdvelop/dbtools => ../dbtools

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/input => ../input
