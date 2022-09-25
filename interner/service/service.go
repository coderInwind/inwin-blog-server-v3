package service

type ServiceGroup struct {
	BlogService
	UserService
	TagService
	GalleryService
	UploadService
	CoverService
}

var ServiceGroupApp = new(ServiceGroup)
