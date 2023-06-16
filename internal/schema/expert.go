package schema

import "admin-server/internal/model"

//request

type ExpertItem struct {
	model.Expert
	SportType ObjIdName `json:"sport_type"`
}

type ExpertIdRequest struct {
	Id int64 `json:"id" validate:"required|gt:0"`
}

type ExpertListRequest struct {
	Id       int64   `json:"id" validate:"gt:0"`
	Name     string  `json:"name" validate:"string|min_len:1|max_len:40"`
	Phone    string  `json:"phone" validate:"string|max_len:11"`
	Statuses []int64 `json:"statuses"  validate:"ints"`
	Page     int     `json:"page"  validate:"gte:-1"`
	PageSize int     `json:"page_size" validate:"gte:-1"`
}

type ExpertListResponse struct {
	Total    int64         `json:"total"`
	Page     int           `json:"page"`
	PageSize int           `json:"page_size"`
	List     []*ExpertItem `json:"list"`
}

type ExpertUpdateRequest struct {
	Id          int64  `json:"id" validate:"required|gt:0"`
	Name        string `json:"name" validate:"string|max_len:40"`
	SportTypeId int64  `json:"sport_type_id" validate:"gt:0"`
}
