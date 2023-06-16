package service

import (
	"admin-server/internal/model"
	"admin-server/internal/schema"
	"admin-server/pkg/utils"
	"gorm.io/gorm"
)

type FeedbackService struct {
	db *gorm.DB
}

func (s *FeedbackService) GetFeedbackList(params schema.FeedbackListRequest) (*schema.FeedbackListResponse, error) {

	dbModel := s.db.Model(&model.Feedback{})

	userWhere := make(map[string]interface{})
	if params.UserName != "" {
		userWhere["name"] = params.UserName
	}

	if params.Phone != "" {
		userWhere["phone"] = params.Phone
	}

	if len(userWhere) > 0 {
		userIds := make([]int64, 0)
		err := s.db.Model(&model.User{}).Where(userWhere).Pluck("id", &userIds).Error
		if err != nil {
			return nil, err
		}
		dbModel.Where("user_id in ?", userIds)
	}

	var total int64
	if err := dbModel.Count(&total).Error; err != nil {
		return nil, err
	}

	var list []schema.FeedbackUser

	err := dbModel.Preload("User", func(db *gorm.DB) *gorm.DB { return db.Model(&model.User{}) }).
		Limit(params.PageSize).Offset(utils.GetOffset(params.PageSize, params.Page)).Order("id desc").Find(&list).Error
	if err != nil {
		return nil, err
	}

	return &schema.FeedbackListResponse{total, params.Page, len(list), list}, nil
}

func (s *FeedbackService) FeedbackUpdate(ids []int64, data map[string]interface{}) error {

	if err := s.db.Model(&model.Feedback{}).Where("id in ?", ids).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}

func NewFeedbackService(db *gorm.DB) *FeedbackService {
	return &FeedbackService{
		db: db,
	}
}
