package service

import (
	"admin-server/internal/model"
	"admin-server/internal/schema"
	"admin-server/pkg/utils"
	"gorm.io/gorm"
)

type CompetitionLinksService struct {
	db *gorm.DB
}

func (s *CompetitionLinksService) GetCompetitionLinks(query interface{}, args ...interface{}) (*schema.CompetitionLinksItem, error) {

	m := schema.CompetitionLinksItem{}

	if err := s.db.Preload("SportType", func(db *gorm.DB) *gorm.DB {
		return db.Model(&model.SportType{})
	}).
		Where(query, args...).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (s *CompetitionLinksService) GetCompetitionLinksList(params schema.CompetitionLinksListRequest) (*schema.CompetitionLinksListResponse, error) {

	dbModel := s.db.Model(&model.CompetitionLinks{}).
		Preload("SportType", func(db *gorm.DB) *gorm.DB {
			return db.Model(&model.SportType{})
		})

	if params.Name != "" {
		dbModel = dbModel.Where("name LIKE ?", "%"+params.Name+"%")
	}

	if params.SportTypeId > 0 {
		dbModel = dbModel.Where("sport_type_id = ?", params.SportTypeId)
	}
	if len(params.Status) > 0 {
		dbModel = dbModel.Where("status in ?", params.Status)
	}

	var total int64
	if err := dbModel.Count(&total).Error; err != nil {
		return nil, err
	}

	var list []*schema.CompetitionLinksItem

	err := dbModel.Limit(params.PageSize).
		Offset(utils.GetOffset(params.PageSize, params.Page)).
		Order("sort desc").
		Order("id desc").
		Find(&list).Error
	if err != nil {
		return nil, err
	}

	return &schema.CompetitionLinksListResponse{total, params.Page, len(list), list}, nil
}

func (s *CompetitionLinksService) CompetitionLinksAdd(params schema.CompetitionLinksAddRequest) error {
	m := model.CompetitionLinks{
		SportTypeId: params.SportTypeId,
		Name:        params.Name,
		Icon:        params.Icon,
		Link:        params.Link,
		Status:      params.Status,
		Sort:        params.Sort,
	}

	if err := s.db.Create(&m).Error; err != nil {
		return err
	}

	return nil
}

func (s *CompetitionLinksService) CompetitionLinksUpdate(params schema.CompetitionLinksUpdateRequest) error {

	data := map[string]interface{}{
		"sport_type_id": params.SportTypeId,
		"name":          params.Name,
		"icon":          params.Icon,
		"link":          params.Link,
		"status":        params.Status,
		"sort":          params.Sort,
	}

	if err := s.db.Model(&model.CompetitionLinks{}).Where("id = ?", params.Id).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}

func (s *CompetitionLinksService) CompetitionLinksDel(id int64) error {

	if err := s.db.Delete(&model.CompetitionLinks{}, id).Error; err != nil {
		return err
	}

	return nil
}

func NewCompetitionLinksService(db *gorm.DB) *CompetitionLinksService {
	return &CompetitionLinksService{
		db: db,
	}
}
