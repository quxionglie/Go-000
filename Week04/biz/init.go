//+build wireinject

package biz

import (
	"github.com/google/wire"
	"week04/internal/dao"
)

//var Provider = wire.NewSet(dao.NewDB, dao.NewUserDao, NewService)

func InitService() (Service, error) {
	wire.Build(NewService, dao.NewUserDao, dao.NewDB)
	return &service{}, nil
}
