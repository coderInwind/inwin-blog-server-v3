package errcode

import "net/http"

type Error struct {
	Code   int
	Msg    string
	Detail string
}

func (e *Error) GetCode() int {
	return e.Code
}

func (e *Error) GetMsg() string {
	return e.Msg
}

func (e *Error) GetDetail() string {
	return e.Detail
}

func NewError(code int, msg string) *Error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

func (e *Error) WithDetail(err string) *Error {
	e.Detail = err
	return e
}

// StatusCode 错误码和http状态的转换
func (e *Error) StatusCode() int {
	switch e.GetCode() {
	case Success.GetCode():
		return http.StatusOK
	case ServerError.GetCode():
		return http.StatusInternalServerError
	case InvalidParams.GetCode():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.GetCode():
		fallthrough
	case UnauthorizedTokenError.GetCode():
		fallthrough
	case UnauthorizedTokenGenerate.GetCode():
		fallthrough
	case UnauthorizedTokenTimeout.GetCode():
		return http.StatusUnauthorized
		//case TooManyRequests.GetCode():
		//	return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError
}
