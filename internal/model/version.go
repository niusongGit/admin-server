package model

import "admin-server/pkg/orm/datatypes"

type Version struct {
	Id            int64           `gorm:"primaryKey;column:id" json:"id"`
	CreatedAt     datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;autoCreateTime" json:"created_at"`
	UpdatedAt     datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"`
	Version       string          `gorm:"type:varchar(20);not null;default:'';comment:版本号 例：v1.0.3，1.0.3" json:"version"`
	VersionNumber int64           `gorm:"type:int;not null;default:0;comment:版本号的int类型 例：103" json:"version_number"`
	UpdateUrl     string          `gorm:"type:varchar(255);not null;default:'';comment:app下载落地页" json:"update_url"`
	UpdateBin     string          `gorm:"type:varchar(255);not null;default:'';comment:app热更包下载地址" json:"update_bin"`
	UpdateLog     string          `gorm:"type:text;not null;comment:更新日志" json:"update_log"` // 更新日志
	Force         int64           `gorm:"type:tinyint;not null;default:0;comment:是否强制更新 0否 1是" json:"force"`
	Status        int64           `gorm:"type:tinyint;not null;default:0;comment:是否开启 0否 1是" json:"status"`
}

func (Version) TableName() string {
	return `version`
}
