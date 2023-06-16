package service

import (
	"admin-server/internal/model"
	"admin-server/internal/schema"
	"admin-server/pkg/utils"
	"gorm.io/gorm"
)

type BlogrollService struct {
	db *gorm.DB
}

func (s *BlogrollService) GetBlogroll(query interface{}, args ...interface{}) (*model.Blogroll, error) {

	m := model.Blogroll{}

	if err := s.db.Where(query, args...).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (s *BlogrollService) GetBlogrollList(params schema.BlogrollListRequest) (*schema.BlogrollListResponse, error) {

	dbModel := s.db.Model(&model.Blogroll{})

	if params.Name != "" {
		dbModel = dbModel.Where("name LIKE ?", "%"+params.Name+"%")
	}

	if params.Category != "" {
		dbModel = dbModel.Where("category LIKE ?", "%"+params.Category+"%")
	}
	if len(params.Status) > 0 {
		dbModel = dbModel.Where("status in ?", params.Status)
	}

	var total int64
	if err := dbModel.Count(&total).Error; err != nil {
		return nil, err
	}

	var list []model.Blogroll

	err := dbModel.Limit(params.PageSize).Offset(utils.GetOffset(params.PageSize, params.Page)).Order("sort desc").Order("id desc").Find(&list).Error
	if err != nil {
		return nil, err
	}

	return &schema.BlogrollListResponse{total, params.Page, len(list), list}, nil
}

func (s *BlogrollService) BlogrollAdd(params schema.BlogrollAddRequest) error {
	m := model.Blogroll{
		Category: params.Category,
		Name:     params.Name,
		Icon:     params.Icon,
		Link:     params.Link,
		Status:   params.Status,
		Sort:     params.Sort,
	}

	if err := s.db.Create(&m).Error; err != nil {
		return err
	}

	return nil
}

func (s *BlogrollService) BlogrollUpdate(params schema.BlogrollUpdateRequest) error {

	data := map[string]interface{}{
		"category": params.Category,
		"name":     params.Name,
		"icon":     params.Icon,
		"link":     params.Link,
		"status":   params.Status,
		"sort":     params.Sort,
	}

	if err := s.db.Model(&model.Blogroll{}).Where("id = ?", params.Id).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}

func (s *BlogrollService) BlogrollDel(id int64) error {

	if err := s.db.Delete(&model.Blogroll{}, id).Error; err != nil {
		return err
	}

	return nil
}

func NewBlogrollService(db *gorm.DB) *BlogrollService {
	return &BlogrollService{
		db: db,
	}
}
