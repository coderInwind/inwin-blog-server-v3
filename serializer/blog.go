package serializer

import (
	"inwind-blog-server-v3/model"
	"time"
)

type Blog struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Src       string    `json:"src"`
	Tag       string    `json:"tag"`
	Overview  string    `json:"overview"`
	Pv        int       `json:"pv"`
	Like      int       `json:"like"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func BuildBlog(item model.Blog) Blog {
	return Blog{
		Id:        item.ID,
		Title:     item.Title,
		Content:   item.Content,
		Src:       item.Src,
		Tag:       item.Tag.Name,
		Overview:  item.Overview,
		Pv:        item.Pv,
		Like:      item.Like,
		CreatedAt: item.BasicModel.CreatedAt,
		UpdatedAt: item.BasicModel.UpdatedAt,
	}
}

func BuildBlogs(items []model.Blog) (blogs []Blog) {
	for _, item := range items {
		blog := BuildBlog(item)
		blogs = append(blogs, blog)
	}
	return blogs
}
