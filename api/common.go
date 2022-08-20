package api

type ErrorMsg struct {
	Msg string `json:"msg"`
}

func ResponseError(err error) ErrorMsg {
	return ErrorMsg{Msg: "参数传递错误"}
}
