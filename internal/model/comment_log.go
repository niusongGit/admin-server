package model

import (
	"admin-server/pkg/orm/datatypes"
)

type CommentLog struct {
	Id            int64           `gorm:"primaryKey;column:id" json:"id"`
	CreatedAt     datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;autoCreateTime" json:"created_at"`
	UpdatedAt     datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"`
	OperationType int64           `gorm:"type:enum('like');not null;comment:操作类型:like-用户点赞" json:"operation_type"` // 操作类型:like-用户点赞
	CommentId     int64           `gorm:"type:bigint;not null;index:idx_comment_id" json:"comment_id"`
	UserId        int64           `gorm:"type:bigint;not null;index:idx_user_id" json:"user_id"`
	Status        int64           `gorm:"type:tinyint;not null;default:0" json:"status"`
}

func (CommentLog) TableName() string {
	return `comment_log`
}
