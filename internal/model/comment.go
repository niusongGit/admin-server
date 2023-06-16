package model

import (
	"admin-server/pkg/orm/datatypes"
)

type Comment struct {
	Id        int64           `gorm:"primaryKey;column:id" json:"id"`
	CreatedAt datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;autoCreateTime" json:"created_at"`
	UpdatedAt datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"`
	ParentId  int64           `gorm:"type:bigint;not null;default:0;comment:父评论ID" json:"parent_id"`                      // 父评论ID
	PostId    int64           `gorm:"type:bigint;not null;default:0;index:idx_post_id;comment:帖子ID" json:"post_id"`       // 帖子ID
	UserId    int64           `gorm:"type:bigint;not null;default:0;index:idx_user_id;comment:发表该评论的用户ID" json:"user_id"` // 发表该评论的用户ID
	ReplyToId int64           `gorm:"type:bigint;not null;default:0;comment:该评论回复的用户ID" json:"reply_to_id"`               // 该评论回复的用户ID
	Level     int64           `gorm:"type:tinyint;default:1;comment:评论层级：1-最顶层；2-第二层，一共只有两层" json:"level"`                // 评论层级：1-最顶层；2-第二层，一共只有两层
	Content   string          `gorm:"type:varchar(255);not null;default:'';comment:评论内容" json:"content"`                  // 评论内容
	IsSticky  int64           `gorm:"type:tinyint;not null;default:0;comment:是否是置顶评论：0-否；1-是" json:"is_sticky"`           // 是否是置顶评论：0-否；1-是
	IsHot     int64           `gorm:"type:tinyint;not null;default:0;comment:是否是热评：0-否；1-是" json:"is_hot"`                // 是否是热评：0-否；1-是
	LikesNum  int64           `gorm:"type:int;not null;default:0;comment:点赞数" json:"likes_num"`                           // 点赞数
	Status    int64           `gorm:"type:tinyint;not null;default:0" json:"status"`
}

func (Comment) TableName() string {
	return `comment`
}
