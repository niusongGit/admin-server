package model

import (
	"admin-server/pkg/orm/datatypes"
)

// UserMember 用户会员表
type UserMember struct {
	Id           int64           `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`
	CreatedAt    datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;autoCreateTime" json:"created_at"`
	UpdatedAt    datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"`
	UserId       int64           `gorm:"column:user_id;type:bigint(20);index:idx_user_id;default:0;NOT NULL" json:"user_id"`
	MemberType   string          `gorm:"column:member_type;type:enum('month','quarter','year');default:month;comment:会员类型: month-月度会员;quater-季度会员;year-年度会员;NOT NULL" json:"member_type"`
	MemberExpire datatypes.XTime `gorm:"column:member_expire;type:timestamp;default:CURRENT_TIMESTAMP;comment:会员过期时间;NOT NULL" json:"member_expire"`
	Status       int             `gorm:"column:status;type:tinyint(4);default:1;comment:会员状态：0-会员已过期；1-会员已开通;NOT NULL" json:"status"`
}

func (m *UserMember) TableName() string {
	return "user_member"
}
