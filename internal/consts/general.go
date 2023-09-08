package consts

const (
	StrAdmin = "admin_%d"
	// 验证码
	PrefixKeyForPhoneVerfiedCode = "phone:verfied:code:"
)

const (
	UserStatusNormal = iota
	UserStatusDisabled
)

const (
	StatusWaiting = iota
	StatusSuccess
	StatusFailed
)

const (
	PredictedResultRed     = "红"
	PredictedResultBlack   = "黑"
	PredictedResultUnknown = "无"
)

const (
	DataDictionaryPropertiesSystem = "system"       // 系统配置
	DataDictionaryPointAndAmount   = "积分与金额"        // 系统配置
	DataDictionaryMemberCategories = "会员"           // 系统配置
	DataDictionarySmsTemplate      = "sms_template" //短信模板
)

const (
	CompetitionStatusNotStarted = iota // 比赛未开始
	CompetitionStatusStarted           // 进行中
	CompetitionStatusEnd               // 已结束
)

const (
	// 订单队列的key（value是订单号和订单回调来源的json字符串）
	RedisKeyForOrder = "order:queue:no:source"
	// redis中订单队列的value模板
	RedisListValueForOrder = "{\"order_no\": \"%s\", \"callback_source\": \"%s\"}"
	RedisSensitiveWord     = "sensitive:word"
	RedisWebSocketWord     = "notify:operation"
)

const (
	WithdrawStatusWaiting = iota // 提现待审核
	WithdrawStatusFailed         // 提现审核失败
	WithdrawStatusSuccess        // 提现审核成功
)

const (
	RootRoleName = "超级管理员"
	SubjectRole  = "role_%d" //Request definition中的sub在casbin_rule表中的格式
)
