package schema

type SystemProperties struct {
	AmountToPoint                            int64  `json:"amount_to_point" validate:"int"`                                // 金额/积分 = 1/10
	GuaranteedPointMultiple                  int64  `json:"guaranteed_point_multiple" validate:"int"`                      // 担保积分是原贴子要求积分的倍数
	RecentCompetitionCount                   int64  `json:"recent_competition_count" validate:"required|int|gt:0"`         // 近期命中率帖子数
	NearCompetitionFinishDisableSendPostTime int64  `json:"near_competition_finish_disable_send_post_time" validate:"int"` // 临近比赛结束禁止发贴时长 单位为分钟 0表示比赛结束前都可以发贴
	NearCompetitionFinishDisableBuyPostTime  int64  `json:"near_competition_finish_disable_buy_post_time" validate:"int"`  // 临近比赛结束禁止购买贴子时长 单位为分钟 0表示比赛结束前都可以购买
	CustomerServiceTelephone                 string `json:"customer_service_telephone" validate:"string"`                  //客服电话
}

type PointAndAmounts []struct {
	Point  int64   `json:"point" validate:"int"`
	Amount float64 `json:"amount" validate:"float"`
}

type PointAndAmountRequest struct {
	PointAndAmounts `json:"point_and_amounts" validate:"slice"`
}

type MemberCategories []struct {
	Name          string  `json:"name"`
	ValidDay      int     `json:"valid_day" validate:"required"`
	PresentPrice  float32 `json:"present_price" validate:"required"`
	OriginalPrice float32 `json:"original_price" validate:"required"`
	Extend        string  `json:"extend"`
}

type MemberCategoriesRequest struct {
	MemberCategories `json:"member_categories" validate:"slice"`
}

type SmsTemplate struct {
	Pwd  string `json:"pwd" validate:"required"`
	Uid  string `json:"uid" validate:"required"`
	Text string `json:"text" validate:"required"`
}
