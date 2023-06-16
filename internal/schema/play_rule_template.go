package schema

import (
	"admin-server/internal/model"
)

//request

type PlayRuleTemplateIdRequest struct {
	Id int64 `json:"id" validate:"required|gt:0"`
}

type PlayRuleTemplateSportType struct {
	model.PlayRuleTemplate
	SportType model.SportType `json:"sport_type"`
}

type PlayRuleTemplateListRequest struct {
	SportTypeId int64   `json:"sport_type_id"`
	Name        string  `json:"name" validate:"string|min_len:1|max_len:20"`
	Type        string  `json:"type" validate:"string|in:score,values,options"`
	Code        string  `json:"code"  validate:"alpha_dash"`
	Status      []int64 `json:"status"  validate:"ints"`
	Page        int     `json:"page"  validate:"gte:-1"`
	PageSize    int     `json:"page_size" validate:"gte:-1"`
}

type PlayRuleTemplateListResponse struct {
	Total    int64                        `json:"total"`
	Page     int                          `json:"page"`
	PageSize int                          `json:"page_size"`
	List     []*PlayRuleTemplateSportType `json:"list"`
}

type PlayRuleTemplateAddRequest struct {
	SportTypeId         int64                      `json:"sport_type_id" validate:"required|gt:0"`
	Name                string                     `json:"name" validate:"required|string|max_len:20"`
	Code                string                     `json:"code"  validate:"required|alpha_dash"`
	Type                string                     `json:"type" validate:"required|string|in:score,values,options"`
	Choices             []*PlayRuleTemplateChoices `json:"choices"`
	PostContentTemplate string                     `json:"post_content_template"`
	Status              int64                      `json:"status"  validate:"in:0,1"`
}

type PlayRuleTemplateChoices struct {
	Code string `json:"code"  validate:"alpha_dash"`
	Name string `json:"name" validate:"string|max_len:50"`
}

type PlayRuleTemplateUpdateRequest struct {
	Id                  int64                      `json:"id" validate:"required|gt:0"`
	SportTypeId         int64                      `json:"sport_type_id" validate:"required|gt:0"`
	Name                string                     `json:"name" validate:"required|string|max_len:20"`
	Code                string                     `json:"code"  validate:"required|alpha_dash"`
	Type                string                     `json:"type" validate:"required|string|in:score,values,options"`
	Choices             []*PlayRuleTemplateChoices `json:"choices"`
	PostContentTemplate string                     `json:"post_content_template"`
	Status              int64                      `json:"status"  validate:"in:0,1"`
}

type PlayRuleTemplateStatusUpdateRequest struct {
	Ids    []int64 `json:"ids" validate:"required|ints|min_len:1"`
	Status int64   `json:"status"  validate:"in:0,1"`
}
