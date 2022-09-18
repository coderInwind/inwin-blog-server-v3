package service

type ServiceGroup struct {
	BlogService
	UserService
	TagService
	GalleryService
	UploadService
}

var ServiceGroupApp = new(ServiceGroup)
