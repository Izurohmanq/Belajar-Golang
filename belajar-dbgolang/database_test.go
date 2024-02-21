package belajardbgolang

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestingEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:tcp(localhost:3306)/golangdb")

	if err != nil {
		panic(err)
	}

	defer db.Close()
	//ginakan db

}
