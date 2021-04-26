package main

import (
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
	ID        uint64    `db:"id"`
	Username  string    `db:"username"`
	Fullname  string    `db:"fullname"`
	Secret    string    `db:"secret"`
	Email     string    `db:"email"`
	DoB       time.Time `db:"dob"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func main() {
	mode := "sqlite3"

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

	//ds := dialect.From("users")

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
		db, _ = NewMySQLDB()
	} else if mode == "sqlite3" {
		db, _ = NewSQLiteDB()
	}
	_, err := db.Exec(insertSQL, args...)
	if err != nil {
		panic(err)
	}
	log.Println("insert successfull")
}

func NewMySQLDB() (*sqlx.DB, error) {
	return sqlx.Connect("mysql", "root:root1234@tcp(127.0.0.1:3306)/gojakarta_dqb?charset=utf8mb4&parseTime=true")
}

func NewSQLiteDB() (*sqlx.DB, error) {
	return sqlx.Connect("sqlite3", "file:gojakarta_dqb.db?cache=shared")
}
