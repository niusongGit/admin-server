package schema

import "admin-server/internal/model"

//request

type UserIdRequest struct {
	Id int64 `json:"id" validate:"required|gt:0"`
}

type UserIdsRequest struct {
	Ids []int64 `json:"ids" validate:"required|ints"`
}

type UserPhoneRequest struct {
	Phone string `json:"phone" validate:"string|max_len:11"`
}

type UserListRequest struct {
	Id       int64   `json:"id" validate:"gt:0"`
	Name     string  `json:"name" validate:"string|min_len:1|max_len:40"`
	Phone    string  `json:"phone" validate:"string|max_len:11"`
	Statuses []int64 `json:"statuses"  validate:"ints"`
	IsAdmin  *int    `json:"is_admin" validate:"in:0,1"`
	Page     int     `json:"page"  validate:"gte:-1"`
	PageSize int     `json:"page_size" validate:"gte:-1"`
}

type UserListResponse struct {
	Total    int64         `json:"total"`
	Page     int           `json:"page"`
	PageSize int           `json:"page_size"`
	List     *[]model.User `json:"list"`
}

type UserUpdateRequest struct {
	Id     int64  `json:"id" validate:"required|gt:0"`
	Name   string `json:"name" validate:"string|max_len:40"`
	Gender string `json:"gender" validate:"in:男,女"`
	Bio    string `json:"bio"`
}

type UserDisableRequest struct {
	Id      int64 `json:"id" validate:"required|gt:0"`
	Disable bool  `json:"disable"`
}

type UserPointChangeRequest struct {
	Id      int64  `json:"id" validate:"required|gt:0"`
	Type    string `json:"type"  validate:"required|in:income,expenditure"`
	Point   int64  `json:"point"  validate:"required"`
	AdminId int64  `json:"admin_id"`
}

type UserPointRecordObject struct {
	Id      int64  `json:"id"`
	Type    string `json:"type"`
	AdminId int64  `json:"admin_id"`
}

type UserBecomeAdminRequest struct {
	Id      int64 `json:"id" validate:"required|gt:0"`
	IsAdmin *int  `json:"is_admin" validate:"required|in:0,1"`
}
