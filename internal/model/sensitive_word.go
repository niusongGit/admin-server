package model

import "admin-server/pkg/orm/datatypes"

type SensitiveWord struct {
	Id        int64           `gorm:"primaryKey;column:id" json:"id"`
	UpdatedAt datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"`
	CreatedAt datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;autoCreateTime" json:"created_at"`
	Status    int64           `gorm:"type:tinyint;default:0;comment:0正常1禁用" json:"status"`
	Word      string          `gorm:"type:varchar(100);index:idx_word,unique,priority:1;not null" json:"word"`
	Type      int             `gorm:"type:tinyint;index:idx_word,unique,priority:1;default:1;comment:1评论2帖子" json:"type"`
	//Level     int             `gorm:"type:tinyint;default:1;comment:敏感等级" json:"level"`
}

func (c SensitiveWord) TableName() string {
	return "sensitive_word"
}
