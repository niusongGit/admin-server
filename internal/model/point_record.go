package model

import (
	"admin-server/pkg/orm/datatypes"
)

type PointRecord struct {
	Id        int64           `gorm:"primaryKey;column:id" json:"id"`
	CreatedAt datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;autoCreateTime" json:"created_at"`
	UpdatedAt datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"`
	UserId    int64           `gorm:"type:bigint;not null;default:0;index:idx_user_id" json:"user_id"`
	Type      string          `gorm:"type:enum('income', 'expenditure');not null;comment:记录类型:income-收入；expenditure-支出" json:"type"`        // 记录类型:income-收入;expenditure-支出
	Object    datatypes.JSON  `gorm:"not null;comment:收入或支出的对象, 如{\"type\": \"user\",\"id\":1}或{\"type\":\"post\",\"id\":1}" json:"object"` // 收入或支出的对象, 如{"type": "user","id":1}或{"type":"post","id":1}
	Point     int64           `gorm:"type:int;not null;default:0;comment:收入或支出的积分" json:"point"`                                            // 收入或支出的积分
	Status    int64           `gorm:"type:tinyint;not null;default:0;comment:0-充值中1-充值失败2-充值成功" json:"status"`
}

func (PointRecord) TableName() string {
	return `point_record`
}
