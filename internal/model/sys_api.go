package model

type SysApi struct {
	Id          int64  `gorm:"primaryKey;column:id" json:"id"`
	Path        string `gorm:"type:varchar(255);not null;default:'';comment:api路径" json:"path"`
	Description string `gorm:"type:varchar(50);not null;default:'';comment:api中文描述" json:"description"`
	ApiGroup    string `gorm:"type:varchar(30);not null;default:'';comment:api分组" json:"api_group"`
	Method      string `gorm:"type:char(10);not null;default:POST;comment:api分组" json:"method"` // 方法:创建POST(默认)|查看GET|更新PUT|删除DELETE
}

func (SysApi) TableName() string {
	return `sys_api`
}
