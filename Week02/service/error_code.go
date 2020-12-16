package service

type BizError struct {
	code string `json:"code"`
	msg  string `json:"msg"`
	args []string
}

var ERROR_CODE_OK = BizError{code: "0000", msg: "处理成功"}
var ERROR_CODE_USER_NOT_FOUNT = BizError{code: "1001", msg: "用户不存在"}
var ERROR_CODE_UNKOWN = BizError{code: "9999", msg: "系统处理异常，请稍后再试"}

func newErrorCode(code string, msg string, args []string) *BizError {
	p := BizError{code: code, msg: msg}
	return &p
}

func (e *BizError) Error() string {
	return e.code + "," + e.msg
}

func (e *BizError) ToApiRes() *ApiRes {
	var apiRes = ApiRes{Code: e.code, Msg: e.msg}
	return &apiRes
}
