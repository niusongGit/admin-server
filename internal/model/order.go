package model

import (
	"admin-server/pkg/orm/datatypes"
)

type Order struct {
	Id                     int64           `gorm:"primaryKey;column:id" json:"id"`
	CreatedAt              datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;autoCreateTime" json:"created_at"`
	UpdatedAt              datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"`
	UserId                 int64           `gorm:"type:bigint;not null;default:0;index:idx_user_id;comment:用户ID" json:"user_id"`                                        // 用户ID
	OrderNo                string          `gorm:"type:varchar(30);not null;default:'';index:idx_order_no,unique;comment:订单号" json:"order_no"`                          // 订单号
	ThirdPartyOrderNo      string          `gorm:"type:varchar(50);not null;default:'';comment:三方订单号" json:"third_party_order_no"`                                      // 三方订单号
	ThirdPartyDatetime     datatypes.XTime `gorm:"type:timestamp;not null;not null;default:CURRENT_TIMESTAMP;comment:三方回调传过来的交易时间" json:"third_party_datetime"`         // 三方回调传过来的交易时间
	TopUpType              string          `gorm:"type:enum('point', 'member');not null;default:'point';comment:充值类型：point-充值积分；member-开通会员" json:"top_up_type"`        // 充值类型:point-充值积分;member-开通会员
	Amount                 float64         `gorm:"type:decimal(20,2);not null;default:0;comment:支付金额" json:"amount"`                                                    // 支付金额
	Properties             datatypes.JSON  `gorm:"type:json;not null;comment:支付类型对应的相关属性:point-{\"point\": 1000}；member-{\"member_type\": \"year\"}" json:"properties"` // 支付类型对应的相关属性:point-{"point": 1000};member-{"member_type": "year"}
	Status                 int64           `gorm:"type:tinyint;not null;default:0;comment:支付状态 0-未支付 1-已支付" json:"status"`                                              // 支付状态:0-充值中;1-充值失败;2-充值成功
	ThirdPartyPlatformCode string          `gorm:"column:third_party_platform_code;type:varchar(10);comment:三方支付平台code;NOT NULL;default:''" json:"third_party_platform_code"`
	CallbackSource         string          `gorm:"column:callback_source;type:varchar(10);comment:订单回调来源：三方平台code、后台管理或agent...;NOT NULL" json:"callback_source"`
}

func (Order) TableName() string {
	return `order`
}
