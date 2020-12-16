// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package biz

import (
	"week04/internal/dao"
)

// Injectors from init.go:

func InitService() (Service, error) {
	db, err := dao.NewDB()
	if err != nil {
		return nil, err
	}
	userDao, err := dao.NewUserDao(db)
	if err != nil {
		return nil, err
	}
	bizService := NewService(userDao)
	return bizService, nil
}
