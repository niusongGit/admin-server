package schema

import (
	"admin-server/internal/model"
)

//request

type PointRecordIdRequest struct {
	Id int64 `json:"id" validate:"required|gt:0"`
}

type PointRecordUser struct {
	model.PointRecord
	User struct {
		Id    int64  `json:"id"`
		Name  string `json:"name"`
		Phone string `json:"phone"`
	} `json:"user"`
}

type PointRecordListRequest struct {
	UserName string  `json:"user_name" validate:"string"`
	Phone    string  `json:"phone" validate:"string"`
	Type     string  `json:"type" validate:"in:income,expenditure"`
	Status   []int64 `json:"status"  validate:"ints"`
	Page     int     `json:"page"  validate:"gte:-1"`
	PageSize int     `json:"page_size" validate:"gte:-1"`
}

type PointRecordListResponse struct {
	Total    int64             `json:"total"`
	Page     int               `json:"page"`
	PageSize int               `json:"page_size"`
	List     []PointRecordUser `json:"list"`
}
type PointRecordStatusUpdateRequest struct {
	Ids    []int64 `json:"ids" validate:"required|ints|min_len:1"`
	Status int64   `json:"status"  validate:"in:0,1,2"`
}
