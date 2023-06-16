package model

import (
	"admin-server/pkg/orm/datatypes"
)

type ExpertApplicationAudit struct {
	Id           int64           `gorm:"primaryKey;column:id" json:"id"`
	CreatedAt    datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;autoCreateTime" json:"created_at"`
	UpdatedAt    datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"`
	UserId       int64           `gorm:"type:bigint;not null;default:0;index:idx_user_id" json:"user_id"`
	Phone        string          `gorm:"type:char(11);not null;default:'';comment:申请人填的手机号" json:"phone"`   // 申请人填的手机号
	Certificates datatypes.JSON  `gorm:"not null;comment:上传的资质和证明：以json array的字符串形式存储" json:"certificates"` // 上传的资质和证明：以json array的字符串形式存储
	Status       int64           `gorm:"type:tinyint;not null;default:0;comment:状态 0 待审核 1 审核成功 2 审核失败" json:"status"`
	Remark       string          `gorm:"type:varchar(255);not null;default:'';comment:审核备注" json:"remark"` // 审核备注
	SportTypeId  int64           `gorm:"type:bigint;not null;default:0;comment:擅长的体育类型的ID" json:"sport_type_id"`
}

func (ExpertApplicationAudit) TableName() string {
	return `expert_application_audit`
}
