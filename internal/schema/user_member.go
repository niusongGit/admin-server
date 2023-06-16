package schema

import "admin-server/internal/model"

//request

type UserMemberIdRequest struct {
	Id int64 `json:"id" validate:"required|gt:0"`
}

type UserMemberListRequest struct {
	Id       int64   `json:"id" validate:"gt:0"`
	UserName string  `json:"user_name" validate:"string|min_len:1|max_len:40"`
	Phone    string  `json:"phone" validate:"string|max_len:11"`
	Statuses []int64 `json:"statuses"  validate:"ints"` //会员状态：0-会员已过期；1-会员已开通
	Page     int     `json:"page"  validate:"gte:-1"`
	PageSize int     `json:"page_size" validate:"gte:-1"`
}

type UserMemberItem struct {
	model.UserMember
	User struct {
		Id    int64  `json:"id"`
		Name  string `json:"name"`
		Phone string `json:"phone"`
	} `json:"user"`
}

type UserMemberListResponse struct {
	Total    int64             `json:"total"`
	Page     int               `json:"page"`
	PageSize int               `json:"page_size"`
	List     []*UserMemberItem `json:"list"`
}
