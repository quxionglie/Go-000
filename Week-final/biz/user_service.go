package biz

import (
	"errors"
	"finaljob/internal/dao"
)

type ApiRes struct {
	Code       string      `json:"code"`
	Msg        string      `json:"msg"`
	BizContent interface{} `json:"bizContent,omitempty"`
}

type Service interface {
	GetUser(username string) *ApiRes
}

type service struct {
	d dao.UserDao
}

func NewService(d dao.UserDao) (s Service) {
	s = &service{
		d: d,
	}
	return
}

func (s *service) GetUser(username string) *ApiRes {
	user, err := s.d.GetUser(username)
	if err != nil {
		if errors.Is(err, dao.NotFound) {
			return ERROR_CODE_USER_NOT_FOUNT.ToApiRes()
		} else {
			return ERROR_CODE_UNKOWN.ToApiRes()
		}
	} else {
		var apiRes = ERROR_CODE_OK.ToApiRes()
		apiRes.BizContent = user
		return apiRes
	}
}
