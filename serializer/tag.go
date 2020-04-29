package serializer

import "myblog/model"

// Tag 标签序列化器
type Tag struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
}

// BuildTag 序列化标签
func BuildTag(item model.Tag) Tag {
	return Tag{
		ID:        item.ID,
		Name:      item.Name,
		CreatedAt: item.CreatedAt.Unix(),
	}
}

// BuildTags 序列化标签列表
func BuildTags(items []model.Tag) (Tags []Tag) {
	for _, item := range items {
		Tag := BuildTag(item)
		Tags = append(Tags, Tag)
	}
	return Tags
}

// BuildTagResponse 序列化标签响应
func BuildTagResponse(Tag model.Tag) Response {
	return Response{
		Data: BuildTag(Tag),
	}
}

// BuildTagsResponse 序列化标签列表响应
func BuildTagsResponse(Tags []model.Tag) Response {
	return Response{
		Data: BuildTags(Tags),
	}
}
