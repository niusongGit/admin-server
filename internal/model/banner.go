package model

type Banner struct {
	Id          int64  `gorm:"primaryKey;column:id" json:"id"`
	SportTypeId int64  `gorm:"type:bigint;not null;default:0;index:idx_sport_type_id" json:"sport_type_id"`
	Name        string `gorm:"type:varchar(40);not null;default:''" json:"name"`
	Icon        string `gorm:"type:varchar(255);not null;default:'';comment:图片地址" json:"icon"`
	Link        string `gorm:"type:varchar(255);not null;default:'';comment:链接地址" json:"link"`
	Status      int64  `gorm:"type:tinyint;not null;default:0" json:"status"`
}

func (Banner) TableName() string {
	return `banner`
}
