package main

import (
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/enix/wal-g/cmd/sqlserver"
)

func main() {
	sqlserver.Execute()
}
