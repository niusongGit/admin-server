package model

import (
	"admin-server/pkg/orm/datatypes"
)

type Post struct {
	Id              int64           `gorm:"primaryKey;column:id" json:"id"`
	CreatedAt       datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;autoCreateTime" json:"created_at"`
	UpdatedAt       datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"`
	UserId          int64           `gorm:"type:bigint;not null;default:0;index:idx_user_id" json:"user_id"`
	CompetitionId   int64           `gorm:"type:bigint;not null;default:0;index:idx_competition_id" json:"competition_id"`
	Title           string          `gorm:"type:varchar(255);not null;default:'';comment:评论内容" json:"title"`
	ExpertAnalysis  string          `gorm:"type:text;not null;comment:专家分析" json:"expert_analysis"` // 专家分析
	Points          float64         `gorm:"type:decimal(20,2);not null;default:0" json:"points"`
	PredictedResult string          `gorm:"type:enum('红', '黑', '无');not null;default:'无';comment:预测结果：无-比赛还未结束，无预测结果" json:"predicted_result"` // 预测结果：无-比赛还未结束，无预测结果
	ReadersNum      int64           `gorm:"type:int;not null;default:0" json:"readers_num"`
	CommentsNum     int64           `gorm:"type:int;not null;default:0" json:"comments_num"`
	LikesNum        int64           `gorm:"type:int;not null;default:0" json:"likes_num"`
	IsEssencePost   int64           `gorm:"type:tinyint;default:0;comment:是否是精华帖：0-否；1-是" json:"is_essence_post"` // 是否是精华帖：0-否；1-是
	PlayRules       datatypes.JSON  `gorm:"type:json;not null;comment:玩法规则（json字符串形式）" json:"play_rules"`         // 玩法规则（json字符串形式）
	Images          datatypes.JSON  `gorm:"type:json;not null;comment:上传的图片" json:"images"`                       // 上传的图片
	IsGuaranteed    int64           `gorm:"type:tinyint;not null;default:0;comment:帖子是否被担保：0-否；1-是" json:"is_guaranteed"`
	IsTop           int64           `gorm:"type:tinyint;not null;default:0;comment:帖子是否置顶：0-否；1-是" json:"is_top"`
	Status          int64           `gorm:"type:tinyint;not null;default:0;comment:状态 0 待审核 1 审核成功 2 审核失败" json:"status"`
	Remark          string          `gorm:"type:varchar(255);not null;default:'';comment:审核备注" json:"remark"` // 审核备注
}

func (Post) TableName() string {
	return `post`
}
