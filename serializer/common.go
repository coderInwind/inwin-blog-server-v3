package serializer

// Response 基础序列化器
type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Msg   string      `json:"msg"`
	Error string      `json:"error"`
}

// 基础的请求响应提示信息
type PromptResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type DataList struct {
	Item  interface{} `json:"item"`
	Total int64       `json:"total"`
}

//BulidListResponse 带有总数的列表构建器
func BuildListResponse(items interface{}, total int64) Response {
	return Response{
		Code: 200,
		Data: DataList{
			Item:  items,
			Total: total,
		},
		Msg: "ok",
	}
}

// 错误提示
func BuildErrorResponse(code int, msg string) PromptResponse {
	return PromptResponse{Code: code, Msg: msg}
}

//操作成功提示
func BuildSuccessResponse() PromptResponse {
	return PromptResponse{Code: 200, Msg: "操作成功"}
}
