package model

type CompetitionType struct {
	Id          int64  `gorm:"primaryKey;column:id" json:"id"`
	SportTypeId int64  `gorm:"type:bigint;not null;default:0;index:idx_sport_type_id" json:"sport_type_id"`
	Name        string `gorm:"type:varchar(40);not null;default:''" json:"name"`
	Icon        string `gorm:"type:varchar(255);not null;default:'';comment:图片地址" json:"icon"`
	Status      int64  `gorm:"type:tinyint;not null;default:0" json:"status"`
	Sort        int64  `gorm:"type:tinyint;not null;default:0;comment:排序" json:"sort"` // 排序
}

func (CompetitionType) TableName() string {
	return `competition_type`
}
