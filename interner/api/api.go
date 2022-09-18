package api

type ApiGroup struct {
	BlogApi
	UserApi
	TagApi
	GalleryApi
	UploadApi
}

var ApiGroupApp = new(ApiGroup)
