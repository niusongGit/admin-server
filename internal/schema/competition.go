package schema

import (
	"admin-server/internal/model"
)

//request

type CompetitionIdRequest struct {
	Id int64 `json:"id" validate:"required|gt:0"`
}

type CompetitionIdsRequest struct {
	Ids []int64 `json:"id" validate:"required|ints"`
}

type CompetitionListRequest struct {
	SportTypeId       int64   `json:"sport_type_id"`
	CompetitionTypeId int64   `json:"competition_type_id"`
	Title             string  `json:"title"  validate:"string"`
	Status            []int64 `json:"status"  validate:"ints"`
	Page              int     `json:"page"  validate:"gte:-1"`
	PageSize          int     `json:"page_size" validate:"gte:-1"`
}

type CompetitionSportTypeCompetitionType struct {
	model.Competition
	SportType       model.SportType       `json:"sport_type"`
	CompetitionType model.CompetitionType `json:"competition_type"`
}

type CompetitionListResponse struct {
	Total    int64                                  `json:"total"`
	Page     int                                    `json:"page"`
	PageSize int                                    `json:"page_size"`
	List     []*CompetitionSportTypeCompetitionType `json:"list"`
}

type CompetitionAddRequest struct {
	SportTypeId       int64  `json:"sport_type_id" validate:"required|gt:0"`
	CompetitionTypeId int64  `json:"competition_type_id" validate:"int"`
	StartTime         int64  `json:"start_time" validate:"required|int|gt:0"`
	EndTime           int64  `json:"end_time" validate:"required|int|gt:0|gt_field:StartTime"`
	Title             string `json:"title" validate:"required|max_len:40"`
	TemplateCode      string `json:"template_code"  validate:"required|alpha_dash"`
	Template          string `json:"template"`
	Status            int64  `json:"status"  validate:"in:0,1,2"` // 比赛状态：0-未开始；1-进行中；2-已结束
}

type CompetitionUpdateRequest struct {
	Id                int64  `json:"id" validate:"required|gt:0"`
	SportTypeId       int64  `json:"sport_type_id" validate:"required|gt:0"`
	CompetitionTypeId int64  `json:"competition_type_id" validate:"int"`
	StartTime         int64  `json:"start_time" validate:"required|int|gt:0"`
	EndTime           int64  `json:"end_time" validate:"required|int|gt:0|gt_field:StartTime"`
	Title             string `json:"title" validate:"required|max_len:40"`
	TemplateCode      string `json:"template_code"  validate:"required|alpha_dash"`
	Template          string `json:"template"`
	Status            int64  `json:"status"  validate:"in:0,1,2"` // 比赛状态：0-未开始；1-进行中；2-已结束
}

type CompetitionStatusUpdateRequest struct {
	Ids    []int64 `json:"ids" validate:"required|ints|min_len:1"`
	Status int64   `json:"status"  validate:"in:0,1,2"`
}

type CompetitionFinishRequest struct {
	Id        int64          `json:"id"`
	PlayRules []PlayRuleItem `json:"play_rules" `
	Template  string         `json:"template"`
}

type PlayRuleItem struct {
	RuleCode   string `json:"rule_code" validate:"alpha_dash"`
	RuleName   string `json:"rule_name"`
	ChoiceCode string `json:"choice_code"  validate:"alpha_dash"`
	ChoiceName string `json:"choice_name"`
}
