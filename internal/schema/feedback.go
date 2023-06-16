package schema

import (
	"admin-server/internal/model"
)

//request

type FeedbackUser struct {
	model.Feedback
	User ObjIdName `json:"user"`
}

type FeedbackListRequest struct {
	UserName string `json:"user_name" validate:"string"`
	Phone    string `json:"phone" validate:"string"`
	Page     int    `json:"page"  validate:"gte:-1"`
	PageSize int    `json:"page_size" validate:"gte:-1"`
}

type FeedbackListResponse struct {
	Total    int64          `json:"total"`
	Page     int            `json:"page"`
	PageSize int            `json:"page_size"`
	List     []FeedbackUser `json:"list"`
}
