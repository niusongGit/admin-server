package model

import (
	"admin-server/pkg/orm/datatypes"
)

type User struct {
	Id               int64           `gorm:"primaryKey;column:id" json:"id"`
	Name             string          `gorm:"type:varchar(20);not null;default:''" json:"name"`
	Avatar           string          `gorm:"type:varchar(255);not null;default:'';comment:头像图片地址" json:"avatar"` // 头像图片地址
	Phone            string          `gorm:"type:char(11);not null;default:'';comment:头像图片地址" json:"phone"`
	Gender           string          `gorm:"type:enum('男', '女');not null;default:'男'" json:"gender"`
	Bio              string          `gorm:"type:varchar(255);not null;default:'';comment:个性签名" json:"bio"`                 // 个性签名
	InvitationCode   string          `gorm:"type:varchar(20);not null;default:'';comment:邀请码" json:"invitation_code"`       // 邀请码
	Follower         int64           `gorm:"type:int;not null;default:0;comment:粉丝数" json:"follower"`                       // 粉丝数
	NameModification int64           `gorm:"type:int;not null;default:1;comment:剩余的昵称修改次数" json:"name_modification"`        // 剩余的昵称修改次数
	FollowerRefId    int64           `gorm:"type:bigint;not null;default:0;comment:用户粉丝关系表的ID" json:"follower_ref_id"`      // 用户粉丝关系表的ID
	Following        int64           `gorm:"type:int;not null;default:0;comment:关注数" json:"following"`                      // 关注数
	FollowingRefId   int64           `gorm:"type:bigint;not null;default:0;comment:用户和关注用户的关系表的ID" json:"following_ref_id"` // 用户和关注用户的关系表的ID
	Posting          int64           `gorm:"type:int;not null;default:0;comment:发帖数" json:"posting"`                        // 发帖数
	Likes            int64           `gorm:"type:int;not null;default:0;comment:获得的点赞数" json:"likes"`                       // 获得的点赞数
	IsExpert         int64           `gorm:"type:tinyint;not null;default:0;comment:是否是专家：0-否；1-是" json:"is_expert"`        // 是否是专家：0-否；1-是
	Status           int64           `gorm:"type:tinyint;not null;default:0;comment:0正常 1禁用" json:"status"`
	Points           float64         `gorm:"type:decimal(20,2);not null;default:0;comment:用户所有积分" json:"points"`                          // 用户所有积分
	InvitedCode      string          `gorm:"type:varchar(20);not null;default:'';comment:填写的别人的邀请码" json:"invited_code"`                  // 填写的别人的邀请码
	DeviceId         string          `gorm:"type:varchar(128);not null;default:'';comment:设备唯一ID" json:"device_id"`                       // 设置唯一ID
	WithdrawAmount   float64         `gorm:"type:decimal(20,2);not null;default:0;comment:可提现金额" json:"withdraw_amount"`                  // 可提现金额
	Platform         string          `gorm:"type:enum('ios', 'android');not null;default:'ios';comment:注册平台：ios；android" json:"platform"` //注册平台：ios；android
	CreatedAt        datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;autoCreateTime" json:"created_at"`
	UpdatedAt        datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"`
	FrozenPoints     float64         `gorm:"type:decimal(20,2);not null;default:0;comment:冻结的积分" json:"frozen_points"` // 冻结的积分
	SportTypeId      int64           `gorm:"type:bigint;not null;default:0;comment:擅长的体育类型的ID" json:"sport_type_id"`
	IsAdmin          int             `gorm:"type:tinyint;not null;default:0;comment:是否是管理员: 0-否；1-是" json:"is_admin"`
	RegisterIp       string          `gorm:"type:varchar(20);not null;default:'';comment:用户注册时的IP" json:"register_ip"`
	LoginIp          string          `gorm:"type:varchar(20);not null;default:'';comment:用户最近登录时的IP" json:"login_ip"`
	LoginAt          datatypes.XTime `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:用户最近登录时间" json:"login_at"`
}

func (User) TableName() string {
	return `user`
}
