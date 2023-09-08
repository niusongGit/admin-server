package model

type Role struct {
	Id   int64  `gorm:"primaryKey;column:id" json:"id"`
	Name string `gorm:"unique;type:varchar(40);not null;default:''" json:"name"`
}

func (Role) TableName() string {
	return `role`
}
