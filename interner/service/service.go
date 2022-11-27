package service

type ServiceGroup struct {
	BlogService
	UserService
	TagService
	GalleryService
	UploadService
	CoverService
	SubmitRecordService
	MenuService
	RoleService
}

var ServiceGroupApp = new(ServiceGroup)
