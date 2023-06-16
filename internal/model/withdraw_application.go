package model

import (
	"admin-server/pkg/orm/datatypes"
)

// 提现申请表

type WithdrawApplication struct {
	Id                int64           `gorm:"primaryKey;column:id" json:"id"`
	CreatedAt         datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;autoCreateTime" json:"created_at"`
	UpdatedAt         datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"`
	UserId            int64           `gorm:"type:bigint;not null;default:0;index:idx_user_id" json:"user_id"`                         // 用户ID
	WithdrawAccountId int64           `gorm:"type:bigint;not null;default:0;index:idx_withdraw_account_id" json:"withdraw_account_id"` // 提现账户ID
	WithdrawAmount    float64         `gorm:"type:decimal(20,2);not null;default:0;comment:提现金额" json:"withdraw_amount"`               // 提现金额
	Status            int64           `gorm:"type:tinyint;not null;default:0;comment:状态 0 待审核  1 审核失败 2 审核成功" json:"status"`
	Remark            string          `gorm:"type:varchar(255);not null;default:'';comment:审核备注" json:"remark"`
}

func (WithdrawApplication) TableName() string {
	return `withdraw_application`
}
