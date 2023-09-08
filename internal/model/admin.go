package model

import (
	"admin-server/pkg/orm/datatypes"
)

type Admin struct {
	Id            int64           `gorm:"primaryKey;column:id" json:"id"`
	AdminName     string          `gorm:"type:varchar(40);not null;index" json:"admin_name"`
	Password      string          `gorm:"type:varchar(255);not null" json:"password"`
	RoleId        int64           `gorm:"type:bigint;not null;default:0;index:idx_role_id;comment:角色ID -1为超级管理" json:"role_id"`
	LastLoginTime datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;autoCreateTime" json:"last_login_time"`
	Status        int64           `gorm:"type:tinyint;default:0;comment:状态" json:"status"`
	CreatedAt     datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;autoCreateTime" json:"created_at"`
	UpdatedAt     datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"`
}

func (Admin) TableName() string {
	return `admin`
}
