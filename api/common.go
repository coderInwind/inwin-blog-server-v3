package api

type ErrorMsg struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

func ResponseError(err error) ErrorMsg {
	return ErrorMsg{Msg: "参数传递错误", Code: 403}
}
