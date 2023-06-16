package schema

import "admin-server/internal/model"

//request

type PaymentWayIdRequest struct {
	Id int64 `json:"id" validate:"required|gt:0"`
}

type PaymentWayListRequest struct {
	PlatformCode string  `json:"platform_code" validate:"string|max_len:15"`
	PlatformName string  `json:"platform_name" validate:"string|max_len:40"`
	Status       []int64 `json:"status"  validate:"ints"`
	Page         int     `json:"page"  validate:"gte:-1"`
	PageSize     int     `json:"page_size" validate:"gte:-1"`
}

type PaymentWayListResponse struct {
	Total    int64              `json:"total"`
	Page     int                `json:"page"`
	PageSize int                `json:"page_size"`
	List     []model.PaymentWay `json:"list"`
}

type PaymentWayAddRequest struct {
	PlatformCode string    `json:"platform_code" validate:"required|max_len:15"`
	PlatformName string    `json:"platform_name" validate:"required|max_len:40"`
	MinBalance   float64   `json:"min_balance" validate:"float"`
	MaxBalance   float64   `json:"max_balance" validate:"float"`
	Amounts      []float64 `json:"amounts" validate:"slice"`
	PayType      string    `json:"pay_type" validate:"required|in:alipay,wechat_pay,bank_card"`
	MerchantNo   string    `json:"merchant_no" validate:"required"`
	SignKey      string    `json:"sign_key" validate:"required"`
	Attach       string    `json:"attach" validate:"required"`
	PayUrl       string    `json:"pay_url" validate:"url"`
	NotifyUrl    string    `json:"notify_url" validate:"url"`
	CallbackUrl  string    `json:"callback_url" validate:"url"`
	Sort         int64     `json:"sort" validate:"int"`
	Status       int64     `json:"status"  validate:"in:0,1"`
}

type PaymentWayUpdateRequest struct {
	Id           int64     `json:"id" validate:"required|gt:0"`
	PlatformCode string    `json:"platform_code" validate:"required|max_len:15"`
	PlatformName string    `json:"platform_name" validate:"required|max_len:40"`
	MinBalance   float64   `json:"min_balance" validate:"float"`
	MaxBalance   float64   `json:"max_balance" validate:"float"`
	Amounts      []float64 `json:"amounts" validate:"slice"`
	PayType      string    `json:"pay_type" validate:"required|in:alipay,wechat_pay,bank_card"`
	MerchantNo   string    `json:"merchant_no" validate:"required"`
	SignKey      string    `json:"sign_key" validate:"required"`
	Attach       string    `json:"attach" validate:"required"`
	PayUrl       string    `json:"pay_url" validate:"url"`
	NotifyUrl    string    `json:"notify_url" validate:"url"`
	CallbackUrl  string    `json:"callback_url" validate:"url"`
	Sort         int64     `json:"sort" validate:"int"`
	Status       int64     `json:"status"  validate:"in:0,1"`
}
