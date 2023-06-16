package service

import (
	"admin-server/internal/model"
	"admin-server/internal/schema"
	"admin-server/pkg/utils"
	"gorm.io/gorm"
)

type CompetitionTypeService struct {
	db *gorm.DB
}

func (s *CompetitionTypeService) GetCompetitionType(query interface{}, args ...interface{}) (*schema.CompetitionTypeSportType, error) {

	m := schema.CompetitionTypeSportType{}

	if err := s.db.Where(query, args...).Preload("SportType").First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (s *CompetitionTypeService) GetCompetitionTypeList(params schema.CompetitionTypeListRequest) (*schema.CompetitionTypeListResponse, error) {

	dbModel := s.db.Model(&model.CompetitionType{})

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

	var list []*schema.CompetitionTypeSportType

	err := dbModel.Preload("SportType").Limit(params.PageSize).Offset(utils.GetOffset(params.PageSize, params.Page)).Order("id desc").Find(&list).Error
	if err != nil {
		return nil, err
	}

	return &schema.CompetitionTypeListResponse{total, params.Page, len(list), list}, nil
}

func (s *CompetitionTypeService) CompetitionTypeAdd(params schema.CompetitionTypeAddRequest) error {
	m := model.CompetitionType{
		SportTypeId: params.SportTypeId,
		Name:        params.Name,
		Icon:        params.Icon,
		Status:      params.Status,
		Sort:        params.Sort,
	}

	if err := s.db.Create(&m).Error; err != nil {
		return err
	}

	return nil
}

func (s *CompetitionTypeService) CompetitionTypeUpdate(params schema.CompetitionTypeUpdateRequest) error {
	data := map[string]interface{}{
		"sport_type_id": params.SportTypeId,
		"name":          params.Name,
		"icon":          params.Icon,
		"status":        params.Status,
		"sort":          params.Sort,
	}

	if err := s.db.Model(&model.CompetitionType{}).Where("id = ?", params.Id).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}

func (s *CompetitionTypeService) CompetitionTypeDel(id int64) error {

	if err := s.db.Delete(&model.CompetitionType{}, id).Error; err != nil {
		return err
	}

	return nil
}

func NewCompetitionTypeService(db *gorm.DB) *CompetitionTypeService {
	return &CompetitionTypeService{
		db: db,
	}
}
