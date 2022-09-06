package api

type ApiGroup struct {
	BlogApi
	UserApi
	TagApi
}

var ApiGroupApp = new(ApiGroup)
