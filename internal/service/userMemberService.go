package service

import (
	"admin-server/internal/model"
	"admin-server/internal/schema"
	"admin-server/pkg/utils"
	"gorm.io/gorm"
)

type UserMemberService struct {
	db *gorm.DB
}

func (s *UserMemberService) GetUserMember(query interface{}, args ...interface{}) (*schema.UserMemberItem, error) {

	m := schema.UserMemberItem{}

	if err := s.db.Where(query, args...).Preload("User",
		func(db *gorm.DB) *gorm.DB {
			return db.Model(&model.User{})
		}).
		First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (s *UserMemberService) GetUserMemberList(params schema.UserMemberListRequest) (*schema.UserMemberListResponse, error) {
	var userIds []int64
	userWhere := make(map[string]interface{})
	if params.UserName != "" {
		userWhere["name"] = params.UserName
	}
	if params.Phone != "" {
		userWhere["phone"] = params.Phone
	}
	if len(userWhere) > 0 {
		err := s.db.Model(&model.User{}).Where(userWhere).Pluck("id", &userIds).Error
		if err != nil {
			return nil, err
		}
		if len(userIds) == 0 {
			return &schema.UserMemberListResponse{0, params.Page, 0, []*schema.UserMemberItem{}}, nil
		}
	}

	var total int64
	umDb := s.db.Model(&model.UserMember{})
	if len(userIds) != 0 {
		umDb.Where("user_id in (?)", userIds)
	}

	if len(params.Statuses) > 0 {
		umDb.Where("status in ?", params.Statuses)
	}
	err := umDb.Count(&total).Error
	var list []*schema.UserMemberItem

	err = umDb.Preload("User", func(db *gorm.DB) *gorm.DB { return db.Model(&model.User{}) }).
		Limit(params.PageSize).Offset(utils.GetOffset(params.PageSize, params.Page)).Order("id desc").
		Find(&list).Error
	if err != nil {
		return nil, err
	}

	return &schema.UserMemberListResponse{total, params.Page, len(list), list}, nil
}

func NewUserMemberService(db *gorm.DB) *UserMemberService {
	return &UserMemberService{
		db: db.Debug(),
	}
}
