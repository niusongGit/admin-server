package schema

import (
	"admin-server/internal/model"
)

type ObjIdName struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

//request

type CommentIdRequest struct {
	Id int64 `json:"id" validate:"required|gt:0"`
}

type CommentIdsRequest struct {
	Ids []int64 `json:"id" validate:"required|ints"`
}

type CommentItem struct {
	model.Comment
	User        ObjIdName  `json:"user"`
	ReplyToUser ObjIdName  `gorm:"foreignKey:reply_to_id;" json:"reply_to_user"`
	Post        model.Post `json:"post"`
}

type CommentListRequest struct {
	UserName string  `json:"user_name" validate:"string"`
	Phone    string  `json:"phone" validate:"string"`
	Content  string  `json:"content"`
	Statuses []int64 `json:"statuses"`
	Level    int64   `json:"level" validate:"in:1,2"`     // 评论层级：1-最顶层；2-第二层，一共只有两层
	IsSticky *int64  `json:"is_sticky" validate:"in:0,1"` // 是否是置顶评论：0-否；1-是
	IsHot    *int64  `json:"is_hot" validate:"in:0,1"`    // 是否是热评：0-否；1-是

	Page     int `json:"page"  validate:"gte:-1"`
	PageSize int `json:"page_size" validate:"gte:-1"`
}

type CommentListResponse struct {
	Total    int64         `json:"total"`
	Page     int           `json:"page"`
	PageSize int           `json:"page_size"`
	List     []CommentItem `json:"list"`
}

type CommentUpdateRequest struct {
	Id      int64  `json:"id" validate:"required|gt:0"`
	Content string `json:"content"`
	Status  int64  `json:"status"  validate:"in:0,1,2"`
}

type CommentPlayRules struct {
	RuleCode   string `json:"rule_code" validate:"alpha"`
	RuleName   string `json:"rule_name" validate:"string|max_len:50"`
	ChoiceCode string `json:"choice_code" validate:"alpha"`
	ChoiceName string `json:"choice_name" validate:"string|max_len:50"`
}

type CommentAuditRequest struct {
	Ids    []int64 `json:"ids" validate:"required|ints|min_len:1"`
	Status int64   `json:"status"  validate:"in:0,1,2"`
}

type CommentStickyRequest struct {
	Ids      []int64 `json:"ids" validate:"required|ints|min_len:1"`
	IsSticky int     `json:"is_sticky" validate:"in:0,1"` // 是否是置顶：0-否；1-是
}

type CommentHotRequest struct {
	Ids   []int64 `json:"ids" validate:"required|ints|min_len:1"`
	IsHot int     `json:"is_hot" validate:"in:0,1"` // 是否是精华帖：0-否；1-是
}
