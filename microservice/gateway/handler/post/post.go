package post

import (
	"forum-gateway/dao"
)

type Api struct {
	Dao dao.Interface
}

func New(i dao.Interface) *Api {
	api := new(Api)
	api.Dao = i
	return api
}

type UpdateInfoRequest struct {
	Id       uint32 `json:"id"`
	Content  string `json:"content"`
	Title    string `json:"title"`
	Category string `json:"category"`
}

type Post struct {
	Content       string `json:"content,omitempty"`
	Title         string `json:"title,omitempty"`
	LastEditTime  string `json:"last_edit_time,omitempty"`
	Category      string `json:"category,omitempty"`
	CreatorId     uint32 `json:"creator_id"`
	CreatorName   string `json:"creator_name,omitempty"`
	CreatorAvatar string `json:"creator_avatar"`
}

type ListResponse struct {
	posts *[]Post
}