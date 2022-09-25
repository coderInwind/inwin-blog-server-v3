package api

type ApiGroup struct {
	BlogApi
	UserApi
	TagApi
	GalleryApi
	UploadApi
	bannerApi
}

var ApiGroupApp = new(ApiGroup)
