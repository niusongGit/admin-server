package schema

import "admin-server/internal/model"

type CompetitionLinksItem struct {
	model.CompetitionLinks
	SportType ObjIdName `json:"sport_type"`
}

//request

type CompetitionLinksIdRequest struct {
	Id int64 `json:"id" validate:"required|gt:0"`
}

type CompetitionLinksListRequest struct {
	SportTypeId int64   `json:"sport_type_id"`
	Name        string  `json:"name" validate:"string|min_len:1|max_len:40"`
	Status      []int64 `json:"status"  validate:"ints"`
	Page        int     `json:"page"  validate:"gte:-1"`
	PageSize    int     `json:"page_size" validate:"gte:-1"`
}

type CompetitionLinksListResponse struct {
	Total    int64                   `json:"total"`
	Page     int                     `json:"page"`
	PageSize int                     `json:"page_size"`
	List     []*CompetitionLinksItem `json:"list"`
}

type CompetitionLinksAddRequest struct {
	SportTypeId int64  `json:"sport_type_id" validate:"required|gt:0"`
	Name        string `json:"name" validate:"required|string|max_len:40"`
	Icon        string `json:"icon" validate:"required|url"`
	Link        string `json:"link" validate:"required|url"`
	Status      int64  `json:"status"  validate:"in:0,1"`
	Sort        uint   `json:"sort" validate:"gt:0"`
}

type CompetitionLinksUpdateRequest struct {
	Id          int64  `json:"id" validate:"required|gt:0"`
	SportTypeId int64  `json:"sport_type_id"`
	Name        string `json:"name" validate:"string|max_len:40"`
	Icon        string `json:"icon" validate:"required|url"`
	Link        string `json:"link" validate:"required|url"`
	Status      int64  `json:"status"  validate:"in:0,1"`
	Sort        uint   `json:"sort"`
}
