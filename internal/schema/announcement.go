package schema

import "admin-server/internal/model"

//request

type AnnouncementIdRequest struct {
	Id int64 `json:"id" validate:"required|gt:0"`
}

type AnnouncementSportType struct {
	model.Announcement
	SportType model.SportType `json:"sport_type"`
}

type AnnouncementListRequest struct {
	SportTypeId []int64 `json:"sport_type_id"`
	Title       string  `json:"title" validate:"max_len:40"`
	Status      []int64 `json:"status"  validate:"ints"`
	Page        int     `json:"page"  validate:"gte:-1"`
	PageSize    int     `json:"page_size" validate:"gte:-1"`
}

type AnnouncementListResponse struct {
	Total    int64                    `json:"total"`
	Page     int                      `json:"page"`
	PageSize int                      `json:"page_size"`
	List     []*AnnouncementSportType `json:"list"`
}

type AnnouncementAddRequest struct {
	SportTypeId int64  `json:"sport_type_id"`
	Title       string `json:"title" validate:"required|max_len:40"`
	Content     string `json:"content" validate:"required"`
	Status      int64  `json:"status"  validate:"in:0,1"`
}

type AnnouncementUpdateRequest struct {
	Id          int64  `json:"id" validate:"required|gt:0"`
	SportTypeId int64  `json:"sport_type_id"`
	Title       string `json:"title" validate:"required|max_len:40"`
	Content     string `json:"content" validate:"required"`
	Status      int64  `json:"status"  validate:"in:0,1"`
}
