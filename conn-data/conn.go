package conndata

import (
	"github.com/posteris/database"
)

type Args struct {
	Type string
	DSN  string
}

type Test struct {
	Name string
	Args
}

func GetTestData() []Test {
	return []Test{
		{
			Name: "Clickhouse",
			Args: Args{
				Type: database.Clickhouse,
				DSN:  "tcp://localhost:9000?database=default",
			},
		},
		{
			Name: "MSSQL",
			Args: Args{
				Type: database.MSSQL,
				DSN:  "sqlserver://sa:Adm1n123@localhost:1433?database=master",
			},
		},
		{
			Name: "MySQL",
			Args: Args{
				Type: database.MySQL,
				DSN:  "root:mysql@tcp(localhost:3306)/mysql?parseTime=true",
			},
		},
		{
			Name: "PostgreSQL",
			Args: Args{
				Type: database.Postgres,
				DSN:  "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable",
			},
		},
		{
			Name: "SQLite",
			Args: Args{
				Type: database.SQLite,
				DSN:  "test.db",
			},
		},
	}
}
