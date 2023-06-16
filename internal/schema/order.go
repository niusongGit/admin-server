package schema

import (
	"admin-server/internal/model"
)

//request

type OrderIdRequest struct {
	Id int64 `json:"id" validate:"required|gt:0"`
}

type OrderUser struct {
	model.Order
	User model.User `json:"user"`
}

type OrderListRequest struct {
	OrderNo           string  `json:"order_no" validate:"string"`
	ThirdPartyOrderNo string  `json:"third_party_order_no" validate:"string"`
	UserName          string  `json:"user_name" validate:"string"`
	Phone             string  `json:"phone" validate:"string"`
	TopUpType         string  `json:"top_up_type"  validate:"in:point,member"`
	Status            []int64 `json:"status"  validate:"ints"`
	Page              int     `json:"page"  validate:"gte:-1"`
	PageSize          int     `json:"page_size" validate:"gte:-1"`
}

type OrderListResponse struct {
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
	List     []OrderUser `json:"list"`
}
type OrderStatusUpdateRequest struct {
	Ids    []int64 `json:"ids" validate:"required|ints|min_len:1"`
	Status int64   `json:"status"  validate:"in:0,1,2"`
}
