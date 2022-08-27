package utils

var MapFlags = map[int]string{
	SUCCESS:       "ok",
	ERROR:         "fail",
	InvalidParams: "请求参数错误",

	ErrorAuthCheckTokenFail:    "token错误",
	ErrorAuthCheckTokenTimeout: "token过期",
	ErrorAuthCheckPasswordFail: "密码错误",
}

func GetMsg(code int) string {
	return MapFlags[code]
}
