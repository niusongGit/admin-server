package schema

import (
	"admin-server/internal/model"
)

//request

type WithdrawAccountIdRequest struct {
	Id int64 `json:"id" validate:"required|gt:0"`
}

type WithdrawAccountUser struct {
	model.WithdrawAccount
	User model.User `json:"user"`
}

type WithdrawAccountListRequest struct {
	Name     string  `json:"name" validate:"string"`
	Number   string  `json:"number"`
	UserName string  `json:"user_name" validate:"string"`
	Phone    string  `json:"phone" validate:"string"`
	Type     string  `json:"type"  validate:"in:alipay,wechat_pay,bank_card"`
	BankType string  `json:"bank_type"`
	Status   []int64 `json:"status"  validate:"ints"`
	Page     int     `json:"page"  validate:"gte:-1"`
	PageSize int     `json:"page_size" validate:"gte:-1"`
}

type WithdrawAccountListResponse struct {
	Total    int64                 `json:"total"`
	Page     int                   `json:"page"`
	PageSize int                   `json:"page_size"`
	List     []WithdrawAccountUser `json:"list"`
}

type WithdrawAccountUpdateRequest struct {
	Id       int64  `json:"id" validate:"required|gt:0"`
	Type     string `json:"type" validate:"required|string|in:alipay,wechat_pay,bank_card"`
	Name     string `json:"name" validate:"required|string"`
	Number   string `json:"number" validate:"required|string|max_len:20"`
	BankType string `json:"bank_type" validate:"required|string"`
	Status   int64  `json:"status"  validate:"in:0,1,2"` // 上传的图片
}
type WithdrawAccountStatusUpdateRequest struct {
	Ids    []int64 `json:"ids" validate:"required|ints|min_len:1"`
	Status int64   `json:"status"  validate:"in:0,1,2"`
}
