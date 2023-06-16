package schema

import "admin-server/internal/model"

//request

type ApplicationAuditListRequest struct {
	Page     int     `json:"page"  validate:"gte:-1"`
	PageSize int     `json:"page_size" validate:"gte:-1"`
	Statuses []int64 `json:"statuses"`
}

type UserExpertApplicationAudit struct {
	model.ExpertApplicationAudit
	User      ObjIdName `json:"user"`
	SportType ObjIdName `json:"sport_type"`
}

type ApplicationAuditListResponse struct {
	Total    int64                         `json:"total"`
	Page     int                           `json:"page"`
	PageSize int                           `json:"page_size"`
	List     []*UserExpertApplicationAudit `json:"list"`
}

type ApplicationAuditUpdateRequest struct {
	Id int64 `json:"id" validate:"required|gt:0"`
}

type ApplicationAuditApproveRequest struct {
	Id         int64  `json:"id" validate:"required|gt:0"`
	IsApproved bool   `json:"is_approved"`
	Remark     string `json:"remark"` // 审核备注

}
