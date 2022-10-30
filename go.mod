module github.com/budimanlai/go-services-ping

go 1.19

replace github.com/budimanlai/go-args => /Users/budimanlai/Documents/projects/go/go-args

replace github.com/budimanlai/go-config => /Users/budimanlai/Documents/projects/go/go-config

require (
	github.com/budimanlai/go-config v0.0.1
	github.com/eqto/dbm v0.14.6
)

require github.com/go-sql-driver/mysql v1.6.0 // indirect
