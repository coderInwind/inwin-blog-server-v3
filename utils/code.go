package utils

const (
	SUCCESS       = 200
	ERROR         = 500
	InvalidParams = 400 //请求参数错误

	ErrorExistUser      = 10002
	ErrorNotExistUser   = 10003
	ErrorFailEncryption = 10006
	ErrorNotCompare     = 10007

	ErrorAuthCheckTokenFail    = 30001 //token 错误
	ErrorAuthCheckTokenTimeout = 30002 //token 过期
	ErrorAuthCheckPasswordFail = 30010
)
