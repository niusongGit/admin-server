package service

import (
	"admin-server/internal/model"
	"admin-server/internal/schema"
	"admin-server/pkg/utils"
	"errors"
	"github.com/goccy/go-json"
	"gorm.io/gorm"
)

type PaymentWayService struct {
	db *gorm.DB
}

func (s *PaymentWayService) GetPaymentWay(query interface{}, args ...interface{}) (*model.PaymentWay, error) {

	m := model.PaymentWay{}

	if err := s.db.Where(query, args...).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (s *PaymentWayService) GetPaymentWayList(params schema.PaymentWayListRequest) (*schema.PaymentWayListResponse, error) {

	dbModel := s.db.Model(&model.PaymentWay{})

	if params.PlatformCode != "" {
		dbModel = dbModel.Where("platform_code = ?", params.PlatformCode)
	}

	if params.PlatformName != "" {
		dbModel = dbModel.Where("platform_name = ?", params.PlatformName)
	}
	if len(params.Status) > 0 {
		dbModel = dbModel.Where("status in ?", params.Status)
	}

	var total int64
	if err := dbModel.Count(&total).Error; err != nil {
		return nil, err
	}

	var list []model.PaymentWay

	err := dbModel.Limit(params.PageSize).Offset(utils.GetOffset(params.PageSize, params.Page)).Order("id desc").Find(&list).Error
	if err != nil {
		return nil, err
	}

	return &schema.PaymentWayListResponse{total, params.Page, len(list), list}, nil
}

func (s *PaymentWayService) PaymentWayAdd(params schema.PaymentWayAddRequest) error {
	amounts, err := json.Marshal(params.Amounts)
	if err != nil {
		return errors.New("amounts数组解析失败：" + err.Error())
	}
	m := model.PaymentWay{
		PlatformCode: params.PlatformCode,
		PlatformName: params.PlatformName,
		MinBalance:   params.MinBalance,
		MaxBalance:   params.MaxBalance,
		Amounts:      amounts,
		PayType:      params.PayType,
		MerchantNo:   params.MerchantNo,
		SignKey:      params.SignKey,
		Attach:       params.Attach,
		PayUrl:       params.PayUrl,
		NotifyUrl:    params.NotifyUrl,
		CallbackUrl:  params.CallbackUrl,
		Sort:         params.Sort,
		Status:       params.Status,
	}

	if err := s.db.Create(&m).Error; err != nil {
		return err
	}

	return nil
}

func (s *PaymentWayService) PaymentWayUpdate(params schema.PaymentWayUpdateRequest) error {
	amounts, err := json.Marshal(params.Amounts)
	if err != nil {
		return errors.New("amounts数组解析失败：" + err.Error())
	}
	data := map[string]interface{}{
		"platform_code": params.PlatformCode,
		"platform_name": params.PlatformName,
		"min_balance":   params.MinBalance,
		"max_balance":   params.MaxBalance,
		"amounts":       amounts,
		"pay_type":      params.PayType,
		"merchant_no":   params.MerchantNo,
		"sign_key":      params.SignKey,
		"attach":        params.Attach,
		"pay_url":       params.PayUrl,
		"notify_url":    params.NotifyUrl,
		"callback_url":  params.CallbackUrl,
		"sort":          params.Sort,
		"status":        params.Status,
	}

	if err := s.db.Model(&model.PaymentWay{}).Where("id = ?", params.Id).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}

func (s *PaymentWayService) PaymentWayDel(id int64) error {

	if err := s.db.Delete(&model.PaymentWay{}, id).Error; err != nil {
		return err
	}

	return nil
}

func NewPaymentWayService(db *gorm.DB) *PaymentWayService {
	return &PaymentWayService{
		db: db,
	}
}
