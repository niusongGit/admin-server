package model

import "admin-server/pkg/orm/datatypes"

type Feedback struct {
	Id         int64           `gorm:"primaryKey;column:id" json:"id"`
	UserId     int64           `gorm:"type:bigint;not null;default:0;index:idx_user_id" json:"user_id"` // 用户ID
	Content    string          `gorm:"type:varchar(300);default:'';column:content" json:"content"`
	Images     datatypes.JSON  `gorm:"type:json;comment:上传的图片" json:"images"` // 上传的图片
	ContactWay string          `gorm:"type:varchar(255);default:'';column:contact_way;comment:联系方式" json:"contact_way"`
	CreatedAt  datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;autoCreateTime" json:"created_at"`
	UpdatedAt  datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"`
}

func (Feedback) TableName() string {
	return `feedback`
}
