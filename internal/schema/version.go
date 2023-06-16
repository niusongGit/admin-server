package schema

import "admin-server/internal/model"

//request

type VersionIdRequest struct {
	Id int64 `json:"id" validate:"required|gt:0"`
}

type VersionListRequest struct {
	Version       string  `json:"version" validate:"string|max_len:20"`
	VersionNumber int64   `json:"version_number" validate:"int"`
	Force         []int64 `json:"force"  validate:"ints"`
	Status        []int64 `json:"status"  validate:"ints"`
	Page          int     `json:"page"  validate:"gte:-1"`
	PageSize      int     `json:"page_size" validate:"gte:-1"`
}

type VersionListResponse struct {
	Total    int64           `json:"total"`
	Page     int             `json:"page"`
	PageSize int             `json:"page_size"`
	List     []model.Version `json:"list"`
}

type VersionAddRequest struct {
	Version       string `json:"version" validate:"required|max_len:20"`
	VersionNumber int64  `json:"version_number" validate:"required"`
	UpdateUrl     string `json:"update_url" validate:"url"`
	UpdateBin     string `json:"update_bin" validate:"url"`
	UpdateLog     string `json:"update_log" validate:"string"`
	Force         int64  `json:"force"  validate:"in:0,1"`
	Status        int64  `json:"status"  validate:"in:0,1"`
}

type VersionUpdateRequest struct {
	Id            int64  `json:"id" validate:"required|gt:0"`
	Version       string `json:"version" validate:"required|max_len:20"`
	VersionNumber int64  `json:"version_number" validate:"required"`
	UpdateUrl     string `json:"update_url" validate:"url"`
	UpdateBin     string `json:"update_bin" validate:"url"`
	UpdateLog     string `json:"update_log" validate:"string"`
	Force         int64  `json:"force"  validate:"in:0,1"`
	Status        int64  `json:"status"  validate:"in:0,1"`
}
