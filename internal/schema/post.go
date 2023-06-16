package schema

import (
	"admin-server/internal/model"
)

//request

type PostIdRequest struct {
	Id int64 `json:"id" validate:"required|gt:0"`
}

type PostIdsRequest struct {
	Ids []int64 `json:"id" validate:"required|ints"`
}

type PostCompetitionUser struct {
	model.Post
	User        model.User                          `json:"user"`
	Competition CompetitionSportTypeCompetitionType `json:"competition"`
}

type PostListRequest struct {
	SportTypeId       int64   `json:"sport_type_id"`
	IsAdPost          int64   `json:"is_ad_post"`
	CompetitionTypeId *int64  `json:"competition_type_id"`
	CompetitionTitle  string  `json:"competition_title" validate:"string"`
	UserName          string  `json:"user_name" validate:"string"`
	Phone             string  `json:"phone" validate:"string"`
	HomeTeam          string  `json:"home_team"  validate:"string"`
	AwayTeam          string  `json:"away_team"  validate:"string"`
	Title             string  `json:"title"  validate:"string"`
	PredictedResult   string  `json:"predicted_result" validate:"in:红,黑,无"` // 预测结果：无-比赛还未结束，无预测结果
	Status            []int64 `json:"status"  validate:"ints"`
	IsGuaranteed      []int64 `json:"is_guaranteed"  validate:"ints"`
	IsTop             []int64 `json:"is_top"  validate:"ints"`
	Page              int     `json:"page"  validate:"gte:-1"`
	PageSize          int     `json:"page_size" validate:"gte:-1"`
}

type PostListResponse struct {
	Total    int64                 `json:"total"`
	Page     int                   `json:"page"`
	PageSize int                   `json:"page_size"`
	List     []PostCompetitionUser `json:"list"`
}

type PostAddRequest struct {
	UserId         int64    `json:"user_id" validate:"required|gt:0"`
	Title          string   `json:"title" validate:"required|string|max_len:255"`
	ExpertAnalysis string   `json:"expert_analysis" validate:"required|string"` // 专家分析float
	IsEssencePost  int64    `json:"is_essence_post" validate:"in:0,1"`          // 是否是精华帖：0-否；1-是
	IsTop          int64    `json:"is_top" validate:"in:0,1"`
	Images         []string `json:"images" validate:"strings"`
}

type PostUpdateAdRequest struct {
	Id             int64    `json:"id" validate:"required|gt:0"`
	UserId         int64    `json:"user_id" validate:"required|gt:0"`
	Title          string   `json:"title" validate:"required|string|max_len:255"`
	ExpertAnalysis string   `json:"expert_analysis" validate:"required|string"` // 专家分析float
	IsEssencePost  int64    `json:"is_essence_post" validate:"in:0,1"`          // 是否是精华帖：0-否；1-是
	IsTop          int64    `json:"is_top" validate:"in:0,1"`
	Images         []string `json:"images" validate:"strings"`
}

type PostUpdateRequest struct {
	Id             int64          `json:"id" validate:"required|gt:0"`
	Title          string         `json:"title" validate:"required|string|max_len:255"`
	ExpertAnalysis string         `json:"expert_analysis" validate:"required|string"` // 专家分析float
	Points         float64        `json:"points" validate:"float"`
	ReadersNum     int64          `json:"readers_num" validate:"int"`
	CommentsNum    int64          `json:"comments_num" validate:"int"`
	LikesNum       int64          `json:"likes_num" validate:"int"`
	IsEssencePost  int64          `json:"is_essence_post" validate:"in:0,1"` // 是否是精华帖：0-否；1-是
	IsTop          int64          `json:"is_top" validate:"in:0,1"`
	PlayRules      *PostPlayRules `json:"play_rules"` // 玩法规则（json字符串形式）
	Images         []string       `json:"images" validate:"strings"`
	IsGuaranteed   int64          `json:"is_guaranteed"  validate:"in:0,1"`
	Status         int64          `json:"status"  validate:"in:0,1,2"`
	Remark         string         `json:"remark"` // 审核备注
	CreatedAt      int64          `json:"created_at" validate:"required|int|gt:0"`
}

type PostPlayRules struct {
	RuleCode   string `json:"rule_code" validate:"alpha_dash"`
	RuleName   string `json:"rule_name" validate:"string|max_len:50"`
	ChoiceCode string `json:"choice_code" validate:"alpha_dash"`
	ChoiceName string `json:"choice_name" validate:"string|max_len:50"`
}

type PostAuditRequest struct {
	Ids    []int64 `json:"ids" validate:"required|ints|min_len:1"`
	Status int64   `json:"status"  validate:"in:0,1,2"`
	Remark string  `json:"remark"` // 审核备注
}

type PostResultRequest struct {
	Id              int64  `json:"id" validate:"required|gt:0"`
	PredictedResult string `json:"predicted_result" validate:"required|string|in:红,黑"` // 预测结果：无-比赛还未结束，无预测结果
}

type PostEssenceRequest struct {
	Ids           []int64 `json:"ids" validate:"required|ints|min_len:1"`
	IsEssencePost int64   `json:"is_essence_post" validate:"in:0,1"` // 是否是精华帖：0-否；1-是
}

type PostTopRequest struct {
	Ids   []int64 `json:"ids" validate:"required|ints|min_len:1"`
	IsTop int64   `json:"is_top" validate:"in:0,1"`
}

type PostLogListRequest struct {
	PostId        int64  `json:"post_id" validate:"required|gt:0"`
	UserName      string `json:"user_name" validate:"string"`
	Phone         string `json:"phone" validate:"string"`
	OperationType string `json:"operation_type" validate:"in:like,favorite,purchase"`
	Page          int    `json:"page"  validate:"gte:-1"`
	PageSize      int    `json:"page_size" validate:"gte:-1"`
}

type PostLogListResponse struct {
	Total    int64          `json:"total"`
	Page     int            `json:"page"`
	PageSize int            `json:"page_size"`
	List     []*PostLogUser `json:"list"`
}

type PostLogUser struct {
	model.PostLog
	User struct {
		Id    int64  `json:"id"`
		Name  string `json:"name"`
		Phone string `json:"phone"`
	} `json:"user"`
}
