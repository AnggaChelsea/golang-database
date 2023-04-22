package database_go

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func GetConnection() *sql.DB {
	dsn := "root:Sayangmamah1.@tcp(localhost:3306)/example_golang_database?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
