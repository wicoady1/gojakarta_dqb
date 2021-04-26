package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/doug-martin/goqu/v9/dialect/sqlite3"
)

type Users struct {
	ID        uint64         `db:"id"`
	Username  string         `db:"username"`
	Fullname  string         `db:"fullname"`
	Address   sql.NullString `db:"address"`
	Secret    string         `db:"secret"`
	Email     string         `db:"email"`
	DoB       time.Time      `db:"dob"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
}

func main() {
	mode := "sqlite3"
	//InsertDemo(mode)
	SelectDemo(mode)
}

func newMySQLDB() (*sqlx.DB, error) {
	return sqlx.Connect("mysql", "root:root1234@tcp(127.0.0.1:3306)/gojakarta_dqb?charset=utf8mb4&parseTime=true")
}

func newSQLiteDB() (*sqlx.DB, error) {
	return sqlx.Connect("sqlite3", "file:sqlite_db/gojakarta_dqb.db?cache=shared")
}

func InsertDemo(mode string) {

	dialect := goqu.Dialect(mode)

	ucokData := Users{
		Username:  "test",
		Fullname:  "Ucok",
		Secret:    "TestPass",
		Email:     "ucok@gmail.com",
		DoB:       time.Date(1993, 6, 6, 0, 0, 0, 0, time.Local),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	/*

		ds := dialect.Insert("users").Cols(
			"username", "fullname", "secret", "email", "dob", "created_at", "updated_at",
		).Vals(
			goqu.Vals{"test", "Ucok", "TestPassword", "ucok@gmail.com", "1993-06-06", time.Now(), time.Now()},
		)
	*/

	ds := dialect.Insert("users").Cols(
		"username", "fullname", "secret", "email", "dob", "created_at", "updated_at",
	).Vals(
		goqu.Vals{
			ucokData.Username,
			ucokData.Fullname,
			ucokData.Secret,
			ucokData.Email,
			ucokData.DoB,
			time.Now(),
			time.Now(),
		},
	)

	insertSQL, args, _ := ds.ToSQL()
	//insertSQL, args, _ := ds.Prepared(true).ToSQL()
	fmt.Println(insertSQL, args)

	//insert into DB
	var db *sqlx.DB
	if mode == "mysql" {
		db, _ = newMySQLDB()
	} else if mode == "sqlite3" {
		db, _ = newSQLiteDB()
	}
	_, err := db.Exec(insertSQL, args...)
	if err != nil {
		panic(err)
	}
	log.Println("insert successful")
}

func SelectDemo(mode string) {
	dialect := goqu.Dialect(mode)
	ds := dialect.From("users")

	query, args, _ := ds.ToSQL()
	fmt.Println(query, args)

	var db *sqlx.DB
	if mode == "mysql" {
		db, _ = newMySQLDB()
	} else if mode == "sqlite3" {
		db, _ = newSQLiteDB()
	}

	var result Users
	err := db.Get(&result, query, args...)
	if err != nil {
		panic(err)
	}

	log.Printf("%+v", result)
}
