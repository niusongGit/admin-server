package model

import (
	"admin-server/pkg/orm/datatypes"
)

type PaymentWay struct {
	Id           int64           `gorm:"primaryKey;column:id" json:"id"`
	CreatedAt    datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;autoCreateTime" json:"created_at"`
	UpdatedAt    datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"`
	PlatformCode string          `gorm:"type:varchar(10);not null;default:'';comment:渠道code，由三方支付提供" json:"platform_code"`
	PlatformName string          `gorm:"type:varchar(36);not null;default:'';comment:渠道名称" json:"platform_name"`
	MinBalance   float64         `gorm:"type:decimal(20,2);not null;default:0;comment:最小金额" json:"min_balance"` //最小金额
	MaxBalance   float64         `gorm:"type:decimal(20,2);not null;default:0;comment:最大金额" json:"max_balance"` //最大金额
	Amounts      datatypes.JSON  `gorm:"type:json;not null;comment:金额数组" json:"amounts"`                        //金额数组
	PayType      string          `gorm:"type:enum('alipay','wechat_pay','bank_card');not null;default:'alipay';comment:支付类型：alipay-支付宝；wechat_pay-微信；bank_card-银行卡" json:"pay_type"`
	MerchantNo   string          `gorm:"type:varchar(30);not null;default:'';comment:支付平台对应的商户号" json:"merchant_no"` //支付平台对应的商户号
	SignKey      string          `gorm:"type:varchar(50);not null;default:'';comment:加密密钥" json:"sign_key"`          //加密key
	Attach       string          `gorm:"type:json;not null;comment:扩展字段可能是对象" json:"attach"`                         //扩展字段可能是对象。
	PayUrl       string          `gorm:"type:varchar(255);not null;default:''" json:"pay_url"`
	NotifyUrl    string          `gorm:"type:varchar(255);not null;default:'';comment:回调地址" json:"notify_url"`
	CallbackUrl  string          `gorm:"type:varchar(255);not null;default:'';comment:支付成功返回地址" json:"callback_url"`
	Sort         int64           `gorm:"type:tinyint;not null;default:0;comment:排序" json:"sort"` // 排序
	Status       int64           `gorm:"type:tinyint;not null;default:0" json:"status"`
}

func (PaymentWay) TableName() string {
	return `payment_way`
}
