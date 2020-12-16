package dao

import (
	"database/sql"
	"errors"
	"fmt"
	xerrors "github.com/pkg/errors"
	"week04/internal/model"
)

var NotFound = errors.New("query no data")

type UserDao interface {
	GetUser(username string) (*model.User, error)
}

type userDao struct {
	db *sql.DB
}

func NewUserDao(db *sql.DB) (d UserDao, err error) {
	d = &userDao{
		db: db,
	}
	return
}

func (dao *userDao) GetUser(username string) (*model.User, error) {
	sqlText := "select Id, Username from user where Username = ?"
	stmt, err := dao.db.Prepare(sqlText)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	user := model.User{}
	err = stmt.QueryRow(username).Scan(&user.Id, &user.Username)
	if err == sql.ErrNoRows {
		return nil, xerrors.Wrap(NotFound, fmt.Sprintf("sql: %s error: %v", sqlText, err))
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}
