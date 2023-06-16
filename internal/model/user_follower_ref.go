package model

type UserFollowerRef struct {
	Id             int64  `gorm:"primaryKey;column:id" json:"id"`
	UserId         int64  `gorm:"type:bigint;not null;default:0;index:idx_user_id" json:"user_id"`
	FollowerUid    int64  `gorm:"type:bigint;not null;default:0;comment:粉丝的用户ID" json:"follower_uid"`             // 粉丝的用户ID
	FollowerName   string `gorm:"type:varchar(40);not null;default:'';comment:粉丝的用户名字" json:"follower_name"`      // 粉丝的用户名字
	FollowerGender string `gorm:"type:enum('男', '女');not null;comment:粉丝的性别" json:"follower_gender"`              // 粉丝的性别
	FollowerAvatar string `gorm:"type:varchar(255);not null;default:'';comment:粉丝的头像图片地址" json:"follower_avatar"` // 粉丝的头像图片地址
}

func (UserFollowerRef) TableName() string {
	return `user_follower_ref`
}
