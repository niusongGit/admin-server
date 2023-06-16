package service

import (
	"admin-server/internal/model"
	"admin-server/internal/schema"
	"admin-server/pkg/orm/datatypes"
	"admin-server/pkg/utils"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
)

type PlayRuleTemplateService struct {
	db *gorm.DB
}

func (s *PlayRuleTemplateService) GetPlayRuleTemplate(query interface{}, args ...interface{}) (*schema.PlayRuleTemplateSportType, error) {

	m := schema.PlayRuleTemplateSportType{}

	if err := s.db.Where(query, args...).Preload("SportType").First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (s *PlayRuleTemplateService) GetPlayRuleTemplateList(params schema.PlayRuleTemplateListRequest) (*schema.PlayRuleTemplateListResponse, error) {

	dbModel := s.db.Model(&model.PlayRuleTemplate{})

	if params.SportTypeId > 0 {
		dbModel = dbModel.Where("sport_type_id = ?", params.SportTypeId)
	}

	if params.Name != "" {
		dbModel = dbModel.Where("name LIKE ?", "%"+params.Name+"%")
	}

	if params.Type != "" {
		dbModel = dbModel.Where("type = ?", params.Type)
	}

	if params.Code != "" {
		dbModel = dbModel.Where("code = ?", params.Code)
	}

	if len(params.Status) > 0 {
		dbModel = dbModel.Where("status in ?", params.Status)
	}

	var total int64
	if err := dbModel.Count(&total).Error; err != nil {
		return nil, err
	}

	var list []*schema.PlayRuleTemplateSportType

	err := dbModel.Preload("SportType").Limit(params.PageSize).Offset(utils.GetOffset(params.PageSize, params.Page)).Order("id desc").Find(&list).Error
	if err != nil {
		return nil, err
	}

	return &schema.PlayRuleTemplateListResponse{total, params.Page, len(list), list}, nil
}

func (s *PlayRuleTemplateService) PlayRuleTemplateAdd(params schema.PlayRuleTemplateAddRequest) error {

	choicesByte, err := json.Marshal(params.Choices)
	if err != nil {
		return errors.New("添加玩法规则模版解析规则数组失败：" + err.Error())
	}
	m := model.PlayRuleTemplate{
		SportTypeId:         params.SportTypeId,
		Name:                params.Name,
		Code:                params.Code,
		Type:                params.Type,
		Choices:             datatypes.JSON(choicesByte),
		PostContentTemplate: params.PostContentTemplate,
		Status:              params.Status,
	}

	if err := s.db.Create(&m).Error; err != nil {
		return err
	}

	return nil
}

func (s *PlayRuleTemplateService) PlayRuleTemplateUpdate(ids []int64, data map[string]interface{}) error {

	if err := s.db.Model(&model.PlayRuleTemplate{}).Where("id in ?", ids).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}

func (s *PlayRuleTemplateService) PlayRuleTemplateDel(id int64) error {

	if err := s.db.Delete(&model.PlayRuleTemplate{}, id).Error; err != nil {
		return err
	}

	return nil
}

func NewPlayRuleTemplateService(db *gorm.DB) *PlayRuleTemplateService {
	return &PlayRuleTemplateService{
		db: db,
	}
}
