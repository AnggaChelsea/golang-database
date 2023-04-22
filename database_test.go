package database_go

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func testEmpty(t *testing.T) {

}

func TestOpenConnect(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/example_table")
	if err != nil {
		panic(err)
	} else {
		defer db.Close()
	}
}
