package serializer

import (
	"inwind-blog-server-v3/common/request"
	"inwind-blog-server-v3/interner/model"
)

func BuildCreateBlogParams(params request.CreateBlog, id uint) model.Blog {
	var tagsSlice []model.Tag

	for _, value := range params.Tags {
		tagsSlice = append(tagsSlice, model.Tag{Id: value})
	}

	return model.Blog{
		BasicModel: model.BasicModel{Id: params.Id},
		Title:      params.Title,
		Content:    params.Content,
		Src:        params.Src,
		Tags:       tagsSlice,
		Overview:   params.Overview,
		AuthorId:   id,
	}
}
