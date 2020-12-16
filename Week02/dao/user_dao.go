package dao

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
}

var ErrQueryNoData = errors.New("query no data")

func GetUser(username string) (*User, error) {
	db, err := getDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmt, err := db.Prepare("select Id, Username from user where Username = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	user := User{}
	err = stmt.QueryRow(username).Scan(&user.Id, &user.Username)
	if err == sql.ErrNoRows {
		return nil, ErrQueryNoData
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func getDb() (*sql.DB, error) {
	db, err := sql.Open("mysql",
		"root:root123@tcp(127.0.0.1:3306)/go")
	return db, err
}
