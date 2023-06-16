package model

type Blogroll struct {
	Id       int64  `gorm:"primaryKey;column:id" json:"id"`
	Category string `gorm:"type:varchar(40);not null;default:''" json:"category"`
	Name     string `gorm:"type:varchar(40);not null;default:''" json:"name"`
	Icon     string `gorm:"type:varchar(255);not null;default:'';comment:图片地址" json:"icon"`
	Link     string `gorm:"type:varchar(255);not null;default:'';comment:链接地址" json:"link"`
	Status   int64  `gorm:"type:tinyint;not null;default:0" json:"status"`
	Sort     uint   `gorm:"type:tinyint;not null;default:0" json:"sort"`
}

func (Blogroll) TableName() string {
	return `blogroll`
}
