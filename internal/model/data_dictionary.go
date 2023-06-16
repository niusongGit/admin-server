package model

import (
	"admin-server/pkg/orm/datatypes"
)

type DataDictionary struct {
	Id         int64           `gorm:"primaryKey;column:id" json:"id"`
	CreatedAt  datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;autoCreateTime" json:"created_at"`
	UpdatedAt  datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"`
	DataType   string          `gorm:"type:varchar(20);not null;default:'';comment:字典类型" json:"data_type"` // 字典类型
	Properties datatypes.JSON  `gorm:"not null;comment:字典类型" json:"properties"`                            // 属性集合
	Status     int64           `gorm:"type:tinyint;not null;default:0" json:"status"`
}

func (DataDictionary) TableName() string {
	return `data_dictionary`
}
