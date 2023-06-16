package service

import (
	"admin-server/internal/model"
	"admin-server/internal/schema"
	"errors"
	"gorm.io/gorm"
)

type SportTypeService struct {
	db *gorm.DB
}

func (s *SportTypeService) GetSportType(query interface{}, args ...interface{}) (*model.SportType, error) {

	m := model.SportType{}

	if err := s.db.Where(query, args...).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (s *SportTypeService) GetSportTypeList(params schema.SportTypeListRequest) ([]model.SportType, error) {

	dbModel := s.db.Model(&model.SportType{})

	if params.Name != "" {
		dbModel = dbModel.Where("name LIKE ?", "%"+params.Name+"%")
	}
	if len(params.Status) > 0 {
		dbModel = dbModel.Where("status in ?", params.Status)
	}

	var list []model.SportType

	err := dbModel.Order("sort desc").Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (s *SportTypeService) SportTypeAdd(params schema.SportTypeAddRequest) error {
	m := model.SportType{}
	s.db.Where("name = ?", params.Name).First(&m)

	if m.Id > 0 {
		return errors.New("已经存在请勿重复添加！！")
	}

	m = model.SportType{
		Name:                                 params.Name,
		PostMaxPoints:                        params.PostMaxPoints,
		PostMinPoints:                        params.PostMinPoints,
		TemplateCode:                         params.TemplateCode,
		Template:                             params.Template,
		CompetitionFinishDisableEditTemplate: params.CompetitionFinishDisableEditTemplate,
		CompetitionAddDisableEditTemplate:    params.CompetitionAddDisableEditTemplate,
		TeamDictionary:                       params.TeamDictionary,
		PostContentTemplate:                  params.PostContentTemplate,
		Sort:                                 params.Sort,
		Status:                               params.Status,
	}
	if err := s.db.Create(&m).Error; err != nil {
		return err
	}

	return nil
}

func (s *SportTypeService) SportTypeUpdate(params schema.SportTypeUpdateRequest) error {
	m := model.SportType{}
	s.db.Where("name = ? and id <> ?", params.Name, params.Id).First(&m)

	if m.Id > 0 {
		return errors.New("已经存在请勿重复编辑！！")
	}

	data := map[string]interface{}{
		"name":            params.Name,
		"post_max_points": params.PostMaxPoints,
		"post_min_points": params.PostMinPoints,
		"template_code":   params.TemplateCode,
		"template":        nil,
		"team_dictionary": nil,
		"competition_finish_disable_edit_template": params.CompetitionFinishDisableEditTemplate,
		"competition_add_disable_edit_template":    params.CompetitionAddDisableEditTemplate,
		"post_content_template":                    params.PostContentTemplate,
		"sort":                                     params.Sort,
		"status":                                   params.Status,
	}

	if params.TeamDictionary != "" {
		data["team_dictionary"] = params.TeamDictionary
	}

	if params.Template != "" {
		data["template"] = params.Template
	}

	if err := s.db.Model(&m).Where("id = ?", params.Id).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}

func (s *SportTypeService) SportTypeDel(id int64) error {

	if err := s.db.Delete(&model.SportType{}, id).Error; err != nil {
		return err
	}

	return nil
}

func NewSportTypeService(db *gorm.DB) *SportTypeService {
	return &SportTypeService{
		db: db,
	}
}
