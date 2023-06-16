package schema

import "admin-server/internal/model"

//request

type BannerIdRequest struct {
	Id int64 `json:"id" validate:"required|gt:0"`
}

type BannerSportType struct {
	model.Banner
	SportType model.SportType `json:"sport_type"`
}

type BannerListRequest struct {
	SportTypeId []int64 `json:"sport_type_id"`
	Name        string  `json:"name" validate:"string|min_len:1|max_len:40"`
	Status      []int64 `json:"status"  validate:"ints"`
	Page        int     `json:"page"  validate:"gte:-1"`
	PageSize    int     `json:"page_size" validate:"gte:-1"`
}

type BannerListResponse struct {
	Total    int64              `json:"total"`
	Page     int                `json:"page"`
	PageSize int                `json:"page_size"`
	List     []*BannerSportType `json:"list"`
}

type BannerAddRequest struct {
	SportTypeId int64  `json:"sport_type_id"`
	Name        string `json:"name" validate:"string|max_len:40"`
	Icon        string `json:"icon" validate:"required|url"`
	Link        string `json:"link" validate:"required|url"`
	Status      int64  `json:"status"  validate:"in:0,1"`
}

type BannerUpdateRequest struct {
	Id          int64  `json:"id" validate:"required|gt:0"`
	SportTypeId int64  `json:"sport_type_id"`
	Name        string `json:"name" validate:"string|max_len:40"`
	Icon        string `json:"icon" validate:"required|url"`
	Link        string `json:"link" validate:"required|url"`
	Status      int64  `json:"status"  validate:"in:0,1"`
}
