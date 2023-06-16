package service

import (
	"admin-server/internal/model"
	"admin-server/internal/schema"
	"admin-server/pkg/utils"
	"gorm.io/gorm"
)

type BannerService struct {
	db *gorm.DB
}

func (s *BannerService) GetBanner(query interface{}, args ...interface{}) (*schema.BannerSportType, error) {

	m := schema.BannerSportType{}

	if err := s.db.Where(query, args...).Preload("SportType").First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (s *BannerService) GetBannerList(params schema.BannerListRequest) (*schema.BannerListResponse, error) {

	dbModel := s.db.Model(&model.Banner{})

	if params.Name != "" {
		dbModel = dbModel.Where("name LIKE ?", "%"+params.Name+"%")
	}
	if len(params.Status) > 0 {
		dbModel = dbModel.Where("status in ?", params.Status)
	}

	if len(params.SportTypeId) > 0 {
		dbModel = dbModel.Where("sport_type_id in ?", params.SportTypeId)
	}

	var total int64
	if err := dbModel.Count(&total).Error; err != nil {
		return nil, err
	}

	var list []*schema.BannerSportType

	err := dbModel.Preload("SportType").Limit(params.PageSize).Offset(utils.GetOffset(params.PageSize, params.Page)).Order("id desc").Find(&list).Error
	if err != nil {
		return nil, err
	}

	return &schema.BannerListResponse{total, params.Page, len(list), list}, nil
}

func (s *BannerService) BannerAdd(params schema.BannerAddRequest) error {
	m := model.Banner{
		SportTypeId: params.SportTypeId,
		Name:        params.Name,
		Icon:        params.Icon,
		Link:        params.Link,
		Status:      params.Status,
	}

	if err := s.db.Create(&m).Error; err != nil {
		return err
	}

	return nil
}

func (s *BannerService) BannerUpdate(params schema.BannerUpdateRequest) error {
	data := map[string]interface{}{
		"sport_type_id": params.SportTypeId,
		"name":          params.Name,
		"icon":          params.Icon,
		"link":          params.Link,
		"status":        params.Status,
	}

	if err := s.db.Model(&model.Banner{}).Where("id = ?", params.Id).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}

func (s *BannerService) BannerDel(id int64) error {

	if err := s.db.Delete(&model.Banner{}, id).Error; err != nil {
		return err
	}

	return nil
}

func NewBannerService(db *gorm.DB) *BannerService {
	return &BannerService{
		db: db,
	}
}
