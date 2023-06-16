package model

// 公告表
type Announcement struct {
	Id          int64  `gorm:"primaryKey;column:id" json:"id"`
	SportTypeId int64  `gorm:"type:bigint;not null;default:0;index:idx_sport_type_id" json:"sport_type_id"`
	Title       string `gorm:"type:varchar(40);not null;default:''" json:"title"`
	Content     string `gorm:"type:text;not null" json:"content"`
	Status      int64  `gorm:"type:tinyint;not null;default:0" json:"status"`
}

func (Announcement) TableName() string {
	return `announcement`
}
