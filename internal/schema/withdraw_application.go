package schema

import (
	"admin-server/internal/model"
)

//request

type WithdrawApplicationIdRequest struct {
	Id int64 `json:"id" validate:"required|gt:0"`
}

type WithdrawApplicationUser struct {
	model.WithdrawApplication
	User            model.User            `json:"user"`
	WithdrawAccount model.WithdrawAccount `json:"withdraw_account"`
}

type WithdrawApplicationListRequest struct {
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

type WithdrawApplicationListResponse struct {
	Total    int64                     `json:"total"`
	Page     int                       `json:"page"`
	PageSize int                       `json:"page_size"`
	List     []WithdrawApplicationUser `json:"list"`
}

type WithdrawApplicationAuditRequest struct {
	Id     int64  `json:"id" validate:"required|gt:0"`
	Status int64  `json:"status"  validate:"in:1,2"`
	Remark string `json:"remark"` // 审核备注
}
