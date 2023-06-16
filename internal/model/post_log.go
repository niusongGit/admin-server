package model

import (
	"admin-server/pkg/orm/datatypes"
)

type PostLog struct {
	Id            int64           `gorm:"primaryKey;column:id" json:"id"`
	CreatedAt     datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;autoCreateTime" json:"created_at"`
	UpdatedAt     datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"`
	OperationType string          `gorm:"type:enum('like', 'favorite', 'purchase');not null;comment:操作类型:like-用户点赞；favorite-用户收藏；purchase-购买方案" json:"operation_type"` // 操作类型:like-用户点赞;favorite-用户收藏;purchase-购买方案
	PostId        int64           `gorm:"type:bigint;not null;default:0;index:idx_post_id" json:"post_id"`
	UserId        int64           `gorm:"type:bigint;not null;default:0;index:idx_user_id" json:"user_id"`
	IsGuaranteed  int64           `gorm:"type:tinyint;not null;default:0;comment:操作类型为purchase时，表示是否购买了担保帖子：0-否；1-是" json:"is_guaranteed"`
	Status        int64           `gorm:"type:tinyint;not null;default:0;comment:操作类型为purchase时，0-未处理 1-已处理" json:"status"`
}

func (PostLog) TableName() string {
	return `post_log`
}
