package model

import (
	"admin-server/pkg/orm/datatypes"
)

// 提现账户表

type WithdrawAccount struct {
	Id        int64           `gorm:"primaryKey;column:id" json:"id"`
	CreatedAt datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;autoCreateTime" json:"created_at"`
	UpdatedAt datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"`
	UserId    int64           `gorm:"type:bigint;not null;default:0;index:idx_user_id" json:"user_id"`                                                                        // 用户ID
	Type      string          `gorm:"type:enum('alipay','wechat_pay','bank_card');not null;default:'alipay';comment:账户类型：alipay-支付宝；wechat_pay-微信；bank_card-银行卡" json:"type"` // 账户类型
	Name      string          `gorm:"type:varchar(10);not null;default:'';comment:账户姓名" json:"name"`                                                                          // 账户姓名
	Number    string          `gorm:"type:varchar(20);not null;default:'';comment:账号或者银行卡号" json:"number"`                                                                    // 账号或者银行卡号
	BankType  string          `gorm:"type:varchar(10);not null;default:'';comment:银行类型" json:"bank_type"`                                                                     // 银行类型
	Status    int64           `gorm:"type:tinyint;not null;default:0" json:"status"`
}

func (WithdrawAccount) TableName() string {
	return `withdraw_account`
}
