package model

import (
	"admin-server/pkg/orm/datatypes"
)

type Competition struct {
	Id                int64           `gorm:"primaryKey;column:id" json:"id"`
	Title             string          `gorm:"type:varchar(128);not null;default:'';comment:标题" json:"title"`
	TemplateCode      string          `gorm:"type:varchar(50);not null;default:'';comment:模板的code" json:"template_code"`
	Template          string          `gorm:"type:json;default:null;comment:模板json" json:"template"`
	SportTypeId       int64           `gorm:"type:bigint;not null;default:0;index:idx_sport_type_id" json:"sport_type_id"`
	CompetitionTypeId int64           `gorm:"type:bigint;not null;default:0;index:idx_competition_type_id" json:"competition_type_id"`
	StartTime         datatypes.XTime `gorm:"type:timestamp;not null" json:"start_time"`
	EndTime           datatypes.XTime `gorm:"type:timestamp;comment:预约结束时间" json:"end_time"`
	PracticalEndTime  datatypes.XTime `gorm:"type:timestamp;comment:实际结束时间" json:"practical_end_time"`
	Status            int64           `gorm:"type:tinyint;not null;default:0;comment:比赛状态：0-未开始；1-进行中；2-已结束" json:"status"`  // 比赛状态：0-未开始；1-进行中；2-已结束
	FinishPlayRules   datatypes.JSON  `gorm:"type:json;default:null;comment:结束比赛最终玩法规则（json字符串形式）" json:"finish_play_rules"` // 比赛结束时选择的最终玩法规则（json字符串形式）
}

func (Competition) TableName() string {
	return `competition`
}
