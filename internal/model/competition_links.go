package model

type CompetitionLinks struct {
	Id          int64  `gorm:"primaryKey;column:id" json:"id"`
	SportTypeId int64  `gorm:"type:bigint;not null;default:0;comment:产品类型id" json:"sport_type_id"`
	Name        string `gorm:"type:varchar(40);not null;default:''" json:"name"`
	Icon        string `gorm:"type:varchar(255);not null;default:'';comment:图片地址" json:"icon"`
	Link        string `gorm:"type:varchar(255);not null;default:'';comment:链接地址" json:"link"`
	Status      int64  `gorm:"type:tinyint;default:0" json:"status"`
	Sort        uint   `gorm:"type:tinyint;default:0" json:"sort"`
}

func (CompetitionLinks) TableName() string {
	return `competition_links`
}
