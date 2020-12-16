package service

import (
	"errors"
	"week02/dao"
)

type ApiRes struct {
	Code       string      `json:"code"`
	Msg        string      `json:"msg"`
	BizContent interface{} `json:"bizContent,omitempty"`
}

func GetUser(username string) *ApiRes {
	user, err := dao.GetUser(username)
	if err != nil {
		if errors.Is(err, dao.ErrQueryNoData) {
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
