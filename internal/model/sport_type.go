package model

type SportType struct {
	Id                                   int64   `gorm:"primaryKey;column:id" json:"id"`
	Name                                 string  `gorm:"type:varchar(40);not null;default:''" json:"name"`
	Sort                                 int64   `gorm:"type:tinyint;not null;default:0;comment:排序" json:"sort"` // 排序
	TemplateCode                         string  `gorm:"type:varchar(50);not null;default:'';comment:模板的code" json:"template_code"`
	Template                             string  `gorm:"type:json;default:null;comment:模板json" json:"template"`
	CompetitionFinishDisableEditTemplate int64   `gorm:"type:tinyint;not null;default:0;comment:结束比赛是否只需要编辑模版 0 否 1 是" json:"competition_finish_disable_edit_template"`
	CompetitionAddDisableEditTemplate    int64   `gorm:"type:tinyint;not null;default:0;comment:新增比赛无需编辑模版 0 否 1 是" json:"competition_add_disable_edit_template"`
	TeamDictionary                       string  `gorm:"type:json;default:null;comment:队伍json字典" json:"team_dictionary"`
	PostMaxPoints                        float64 `gorm:"type:decimal(20,2);not null;default:0;comment:帖子最大积分" json:"post_max_points"`
	PostMinPoints                        float64 `gorm:"type:decimal(20,2);not null;default:0;comment:帖子最小积分" json:"post_min_points"`
	PostContentTemplate                  string  `gorm:"type:text;comment:贴子内容模版" json:"post_content_template"` // 贴子内容模版
	Status                               int64   `gorm:"type:tinyint;not null;default:0" json:"status"`
}

func (SportType) TableName() string {
	return `sport_type`
}
