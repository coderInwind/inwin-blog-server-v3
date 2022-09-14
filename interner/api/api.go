package api

type ApiGroup struct {
	BlogApi
	UserApi
	TagApi
	GalleryApi
}

var ApiGroupApp = new(ApiGroup)
