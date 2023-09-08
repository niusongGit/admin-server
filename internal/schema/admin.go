package schema

import "admin-server/internal/model"

//request

type AdminRequest struct {
	Id        int64  `json:"id" validate:"gt:0"`
	AdminName string `json:"admin_name" validate:"required|string|min_len:4|max_len:10"`
	Password  string `json:"password" validate:"required|min_len:6|max_len:14"`
	RoleId    int64  `json:"role_id" validate:"gte:-1"`
}

type AdminListRequest struct {
	AdminName string  `json:"admin_name" validate:"string|min_len:1|max_len:40"`
	RoleIds   []int64 `json:"role_ids"  validate:"ints"`
	Page      int     `json:"page"  validate:"gte:-1"`
	PageSize  int     `json:"page_size" validate:"gte:-1"`
}

type AdminListResponse struct {
	Total    int64        `json:"total"`
	Page     int          `json:"page"`
	PageSize int          `json:"page_size"`
	List     []*AdminList `json:"list"`
}

type AdminList struct {
	model.Admin
	Role model.Role `json:"role"`
}

type AdminChangePasswordRequest struct {
	OldPassword     string `json:"old_password" validate:"required|min_len:6|max_len:14"`
	NewPassword     string `json:"new_password" validate:"required|min_len:6|max_len:14"`
	ConfirmPassword string `json:"confirm_password" validate:"required|min_len:6|max_len:14"`
}
