package model

import (
	"admin-server/pkg/orm/datatypes"
)

type PlayRuleTemplate struct {
	Id                  int64           `gorm:"primaryKey;column:id" json:"id"`
	CreatedAt           datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;autoCreateTime" json:"created_at"`
	UpdatedAt           datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"`
	SportTypeId         int64           `gorm:"type:bigint;not null;default:0;index:idx_sport_type_id" json:"sport_type_id"`
	Name                string          `gorm:"type:varchar(20);not null;default:'';comment:玩法规则名字" json:"name"`                               // 玩法规则名字
	Code                string          `gorm:"type:varchar(20);not null;default:'';comment:玩法规则名字的code" json:"code"`                          // 玩法规则名字的code
	Type                string          `gorm:"type:enum('score', 'values', 'options');not null;default:'options';comment:玩法规则类型" json:"type"` // 玩法规则类型
	Choices             datatypes.JSON  `gorm:"not null;comment:玩法选择" json:"choices"`                                                          // 玩法选择
	PostContentTemplate string          `gorm:"type:text;comment:贴子内容模版" json:"post_content_template"`                                         // 贴子内容模版
	Status              int64           `gorm:"type:tinyint;not null;default:0;comment:玩法模版状态：0-禁用；1-启用" json:"status"`                        // 玩法模版状态：0-禁用；1-启用
}

func (PlayRuleTemplate) TableName() string {
	return `play_rule_template`
}
