package model

type UserFollowingRef struct {
	Id              int64  `gorm:"primaryKey;column:id" json:"id"`
	UserId          int64  `gorm:"type:bigint;not null;default:0;index:idx_user_id" json:"user_id"`
	FollowingUid    int64  `gorm:"type:bigint;not null;default:0;comment:关注用户的ID" json:"following_uid"`               // 关注用户的ID
	FollowingName   string `gorm:"type:varchar(40);not null;default:'';comment:关注用户的名字" json:"following_name"`        // 关注用户的名字
	FollowingGender string `gorm:"type:enum('男', '女');not null;comment:关注用户的性别" json:"following_gender"`              // 关注用户的性别
	FollowingAvatar string `gorm:"type:varchar(255);not null;default:'';comment:关注用户的头像图片地址" json:"following_avatar"` // 关注用户的头像图片地址
}

func (UserFollowingRef) TableName() string {
	return `user_following_ref`
}
