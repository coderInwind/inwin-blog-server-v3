package serializer

import (
	"inwind-blog-server-v3/common/request"
	"inwind-blog-server-v3/interner/model"
)

//type Blog struct {
//	Id        uint      `json:"id"`
//	Title     string    `json:"title"`
//	Content   string    `json:"content"`
//	Src       string    `json:"src"`
//	Overview  string    `json:"overview"`
//	Pv        int       `json:"pv"`
//	Like      int       `json:"like"`
//	CreatedAt time.Time `json:"createdAt"`
//	UpdatedAt time.Time `json:"updatedAt"`
//}

//func BuildBlogs(list []model.Blog) (blogs []Blog) {
//	for _, item := range list {
//		blog := BuildBlog(item)
//		blogs = append(blogs, blog)
//	}
//	return blogs
//}
//
//func BuildBlog(item model.Blog) Blog {
//	return Blog{
//		Id:        item.ID,
//		Title:     item.Title,
//		Content:   item.Content,
//		Src:       item.Src,
//		Overview:  item.Overview,
//		Pv:        item.Pv,
//		Like:      item.Like,
//		CreatedAt: item.BasicModel.CreatedAt,
//		UpdatedAt: item.BasicModel.UpdatedAt,
//	}
//}

func BuildCreateBlogParams(params request.CreateBlog) model.Blog {
	var tagsSlice []model.Tag

	for _, value := range params.Tags {
		tagsSlice = append(tagsSlice, model.Tag{Id: value})
	}

	return model.Blog{
		Title:    params.Title,
		Content:  params.Content,
		Src:      params.Src,
		Tags:     tagsSlice,
		Overview: params.Overview,
	}
}

func BuildEditBlogParams(params request.EditBlog, allTags []model.Tag) model.Blog {
	var tagsSlice []model.Tag

	for _, value1 := range allTags {
		for _, value2 := range params.Tags {
			if value1.Id == value2 {
				tagsSlice = append(tagsSlice, model.Tag{Id: value2, Name: value1.Name})
			}
		}
	}

	return model.Blog{
		BasicModel: model.BasicModel{ID: params.Id},
		Title:      params.Title,
		Content:    params.Content,
		Src:        params.Src,
		Tags:       tagsSlice,
		Overview:   params.Overview,
	}
}
