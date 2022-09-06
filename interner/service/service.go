package service

type ServiceGroup struct {
	BlogService
	UserService
	TagService
}

var ServiceGroupApp = new(ServiceGroup)
