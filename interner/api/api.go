package api

type ApiGroup struct {
	BlogApi
	UserApi
	TagApi
	GalleryApi
	UploadApi
	bannerApi
	MenuApi
	RoleApi
}

var ApiGroupApp = new(ApiGroup)
