package service

import (
	"admin-server/internal/consts"
	"admin-server/internal/model"
	"admin-server/internal/schema"
	"admin-server/pkg/orm/datatypes"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
)

type SystemService struct {
	db *gorm.DB
}

func (s *SystemService) GetSystem() (*schema.SystemProperties, error) {

	m := model.DataDictionary{}

	if err := s.db.Where(model.DataDictionary{DataType: consts.DataDictionaryPropertiesSystem}).First(&m).Error; err != nil {
		return nil, err
	}

	var properties schema.SystemProperties
	byteArr, _ := m.Properties.MarshalJSON()
	if err := json.Unmarshal(byteArr, &properties); err != nil {
		return nil, err
	}

	return &properties, nil
}

func (s *SystemService) SystemUpdate(params schema.SystemProperties) error {
	propertiesByte, err := json.Marshal(params)
	if err != nil {
		return errors.New("修改系统配置数组json序列化失败：" + err.Error())
	}

	if err := s.db.Model(&model.DataDictionary{}).Where("data_type= ?", consts.DataDictionaryPropertiesSystem).Update("properties", datatypes.JSON(propertiesByte)).Error; err != nil {
		return err
	}

	return nil
}

func (s *SystemService) GetPointAndAmount() (*schema.PointAndAmounts, error) {

	m := model.DataDictionary{}

	if err := s.db.Where(model.DataDictionary{DataType: consts.DataDictionaryPointAndAmount}).First(&m).Error; err != nil {
		return nil, err
	}

	var pointAndAmounts schema.PointAndAmounts
	byteArr, _ := m.Properties.MarshalJSON()
	if err := json.Unmarshal(byteArr, &pointAndAmounts); err != nil {
		return nil, err
	}

	return &pointAndAmounts, nil
}

func (s *SystemService) PointAndAmountUpdate(params schema.PointAndAmountRequest) error {

	pointAndAmountsByte, err := json.Marshal(params.PointAndAmounts)
	if err != nil {
		return errors.New("修改系统配置数组json序列化失败：" + err.Error())
	}

	if err := s.db.Model(&model.DataDictionary{}).Where("data_type= ?", consts.DataDictionaryPointAndAmount).Update("properties", datatypes.JSON(pointAndAmountsByte)).Error; err != nil {
		return err
	}

	return nil
}

func (s *SystemService) GetMemberCategories() (*schema.MemberCategories, error) {

	m := model.DataDictionary{}

	if err := s.db.Where(model.DataDictionary{DataType: consts.DataDictionaryMemberCategories}).First(&m).Error; err != nil {
		return nil, err
	}

	var memberCategories schema.MemberCategories
	byteArr, _ := m.Properties.MarshalJSON()
	if err := json.Unmarshal(byteArr, &memberCategories); err != nil {
		return nil, err
	}

	return &memberCategories, nil
}

func (s *SystemService) MemberCategoriesUpdate(params schema.MemberCategoriesRequest) error {

	memberCategoriesByte, err := json.Marshal(params.MemberCategories)
	if err != nil {
		return errors.New("修改系统配置数组json序列化失败：" + err.Error())
	}

	if err = s.db.Model(&model.DataDictionary{}).Where("data_type= ?", consts.DataDictionaryMemberCategories).Update("properties", datatypes.JSON(memberCategoriesByte)).Error; err != nil {
		return err
	}

	return nil
}

func (s *SystemService) GetSmsTemplate() (*schema.SmsTemplate, error) {

	m := model.DataDictionary{}

	if err := s.db.Where(model.DataDictionary{DataType: consts.DataDictionarySmsTemplate}).First(&m).Error; err != nil {
		return nil, err
	}

	var properties schema.SmsTemplate
	byteArr, _ := m.Properties.MarshalJSON()
	if err := json.Unmarshal(byteArr, &properties); err != nil {
		return nil, err
	}

	return &properties, nil
}

func (s *SystemService) SmsTemplateUpdate(params schema.SmsTemplate) error {
	propertiesByte, err := json.Marshal(params)
	if err != nil {
		return errors.New("修改系统配置数组json序列化失败：" + err.Error())
	}

	if err := s.db.Model(&model.DataDictionary{}).Where("data_type= ?", consts.DataDictionarySmsTemplate).Update("properties", datatypes.JSON(propertiesByte)).Error; err != nil {
		return err
	}

	return nil
}

func NewSystemService(db *gorm.DB) *SystemService {
	return &SystemService{
		db: db,
	}
}
