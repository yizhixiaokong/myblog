package serializer

import "myblog/model"

// Blog 博客序列化器
type Blog struct {
	ID        uint   `json:"id"`
	Author    string `json:"author"`
	Title     string `json:"title"`
	Details   string `json:"details"`
	CreatedAt int64  `json:"created_at"`
}

// BuildBlog 序列化博客
func BuildBlog(item model.Blog) Blog {
	return Blog{
		ID:        item.ID,
		Author:    item.Author,
		Title:     item.Title,
		Details:   item.Details,
		CreatedAt: item.CreatedAt.Unix(),
	}
}

// BuildBlogs 序列化博客列表
func BuildBlogs(items []model.Blog) (blogs []Blog) {
	for _, item := range items {
		blog := BuildBlog(item)
		blogs = append(blogs, blog)
	}
	return blogs
}

// BuildBlogResponse 序列化博客响应
func BuildBlogResponse(blog model.Blog) Response {
	return Response{
		Data: BuildBlog(blog),
	}
}

// BuildBlogsResponse 序列化博客列表响应
func BuildBlogsResponse(blogs []model.Blog) Response {
	return Response{
		Data: BuildBlogs(blogs),
	}
}
