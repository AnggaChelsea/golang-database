package database_go

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("error closin database: %v", err)
		}
	}()
	ctx := context.Background()
	script := " INSERT INTO example_table(id, name, age) VALUES(5, 'angga5', 24)"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}
	fmt.Println("success insert data")
}

func TestExecSqlParams(t *testing.T) {
	db := GetConnection()
	username := "aga'; #"
	password := "aga"
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("error closin database: %v", err)
		}
	}()
	ctx := context.Background()
	script := " INSERT INTO users(username, password) VALUES(?, ?)"
	_, err := db.ExecContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	fmt.Println("success new users data")
}

func TestQuesrSql(t *testing.T) {
	//jikalau mau select query di rekomendasikan manggil satu satu by fiel
	db := GetConnection()
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("error closin database: %v", err)
		}
	}()
	ctx := context.Background()
	script := "SELECT id, name, age FROM example_table" //"SELECT  * FROM example_table"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	fmt.Println("success Get data", rows)
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			panic(err)
		}
		fmt.Println("id", id)
		fmt.Println("name", name)
		fmt.Println("age", age)
	}
}

// jikalau mau select query di rekomendasikan manggil satu satu by fiel
func TestQuesrComplexSql(t *testing.T) {
	db := GetConnection()
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("error closin database: %v", err)
		}
	}()
	ctx := context.Background()
	script := "SELECT id, name, age, email, balance, rating, created_at, married FROM example_table" //"SELECT  * FROM example_table"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	fmt.Println("success Get data", rows)
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var email string
		var age int
		var balance int32
		var rating int64
		var created_at time.Time
		var married sql.NullBool
		err := rows.Scan(&id, &name, &age, &email, &balance, &rating, &created_at, &married)
		if err != nil {
			panic(err)
		}
		fmt.Println("=============")
		fmt.Println("id", id)
		fmt.Println("name", name)
		fmt.Println("age", age)
		fmt.Println("email", email)
		fmt.Println("balance", balance)
		fmt.Println("rating", rating)
		fmt.Println("created_at", created_at)
		if married.Valid {
			fmt.Println("married", married.Bool)
		} else {
			fmt.Println("married", nil)
		}

	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("error closin database: %v", err)
		}
	}()
	ctx := context.Background()

	username := "admin'; #"
	password := "admin"
	//script := "SELECT username from users where username = '" + username + "' and password = '" + password + "' limit 1 "
	//yang di atas bahasa disana belum handle sql injection

	script := "SELECT username from users where username = ? and password  = ? limit 1 "
	//jadi untuk menghindari sql injection pake ? untuk select data

	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	fmt.Println("success Get data", rows)
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("success logn", username)
	} else {
		fmt.Println("gagal login")
	}
}
