package service

import (
	"admin-server/internal/model"
	"admin-server/internal/schema"
	"admin-server/pkg/utils"
	"gorm.io/gorm"
)

type PointRecordService struct {
	db *gorm.DB
}

func (s *PointRecordService) GetPointRecord(query interface{}, args ...interface{}) (*schema.PointRecordUser, error) {

	m := schema.PointRecordUser{}

	err := s.db.Model(model.PointRecord{}).
		Preload("User", func(db *gorm.DB) *gorm.DB { return db.Model(&model.User{}) }).
		Where(query, args...).First(&m).Error
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (s *PointRecordService) GetPointRecordList(params schema.PointRecordListRequest) (*schema.PointRecordListResponse, error) {

	dbModel := s.db.Model(&model.PointRecord{})

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
	if len(params.Status) > 0 {
		dbModel.Where("status in ?", params.Status)
	}
	if params.Type != "" {
		dbModel.Where("type = ?", params.Type)
	}

	var total int64
	if err := dbModel.Count(&total).Error; err != nil {
		return nil, err
	}

	var list []schema.PointRecordUser

	err := dbModel.Preload("User", func(db *gorm.DB) *gorm.DB { return db.Model(&model.User{}) }).
		Limit(params.PageSize).Offset(utils.GetOffset(params.PageSize, params.Page)).Order("id desc").Find(&list).Error
	if err != nil {
		return nil, err
	}

	return &schema.PointRecordListResponse{total, params.Page, len(list), list}, nil
}

func (s *PointRecordService) PointRecordUpdate(ids []int64, data map[string]interface{}) error {

	if err := s.db.Model(&model.PointRecord{}).Where("id in ?", ids).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}

func NewPointRecordService(db *gorm.DB) *PointRecordService {
	return &PointRecordService{
		db: db,
	}
}
