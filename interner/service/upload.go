package service

import "inwind-blog-server-v3/utils"

type UploadService struct {
}

func (UploadService) GetUploadService() string {

	token := utils.NewQiniu().UpLoadToken()

	return token
}
