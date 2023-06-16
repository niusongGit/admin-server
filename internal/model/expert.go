package model

type Expert struct {
	Id             int64   `json:"id"`
	Name           string  `json:"name"`
	SportTypeId    int64   `json:"sport_type_id"`
	RecentHitRatio float32 `gorm:"-" json:"recentHitRatio"` //近期命中率的百分比
	Follower       int64   `json:"follower"`                // 粉丝数
	Following      int64   `json:"following"`               // 关注数
	Posting        int64   `json:"posting"`                 // 发帖数
	Likes          int64   `json:"likes"`                   // 获得的点赞数
	Points         float64 `json:"points"`                  // 用户所有积分
	WithdrawAmount float64 `json:"withdraw_amount"`         // 可提现金额
}
