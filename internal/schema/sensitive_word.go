package schema

import "admin-server/internal/model"

//request

type SensitiveWordIdRequest struct {
	Id int64 `json:"id" validate:"required|gt:0"`
}

type SensitiveWordListRequest struct {
	Type     int     `json:"type"`
	Word     string  `json:"word"`
	Statuses []int64 `json:"statuses"  validate:"ints"`
	Page     int     `json:"page"  validate:"gte:-1"`
	PageSize int     `json:"page_size" validate:"gte:-1"`
}

type SensitiveWordListResponse struct {
	Total    int64                  `json:"total"`
	Page     int                    `json:"page"`
	PageSize int                    `json:"page_size"`
	List     []*model.SensitiveWord `json:"list"`
}

type SensitiveWordAddRequest struct {
	Word string `json:"word" validate:"required"`
	Type int    `json:"type" validate:"in:1,2"`
}

type SensitiveWordUpdateRequest struct {
	Id     int64  `json:"id" validate:"required|gt:0"`
	Word   string `json:"word"`
	Type   int    `json:"type" validate:"in:1,2"`
	Status *int64 `json:"status"  validate:"in:0,1"`
}
