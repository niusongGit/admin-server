package service

import (
	"admin-server/internal/model"
	"admin-server/internal/schema"
	"admin-server/pkg/utils"
	"gorm.io/gorm"
)

type WithdrawAccountService struct {
	db *gorm.DB
}

func (s *WithdrawAccountService) GetWithdrawAccount(query interface{}, args ...interface{}) (*schema.WithdrawAccountUser, error) {

	m := schema.WithdrawAccountUser{}

	if err := s.db.Model(model.WithdrawAccount{}).Preload("User").Where(query, args...).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (s *WithdrawAccountService) GetWithdrawAccountList(params schema.WithdrawAccountListRequest) (*schema.WithdrawAccountListResponse, error) {

	dbModel := s.db.Model(&model.WithdrawAccount{})

	userWhere := make(map[string]interface{})
	if params.UserName != "" {
		userWhere["name"] = params.UserName
	}

	if params.Phone != "" {
		userWhere["phone"] = params.Phone
	}

	if len(userWhere) > 0 {
		userIds := make([]int64, 0)
		s.db.Model(&model.User{}).Where(userWhere).Pluck("id", &userIds)
		dbModel = dbModel.Where("user_id in ?", userIds)
	}

	if params.Name != "" {
		dbModel = dbModel.Where("name LIKE ?", "%"+params.Name+"%")
	}

	if params.Number != "" {
		dbModel = dbModel.Where("number  LIKE ?", "%"+params.Number+"%")
	}

	if params.Type != "" {
		dbModel = dbModel.Where("type = ?", params.Type)
	}

	if params.BankType != "" {
		dbModel = dbModel.Where("bank_type = ?", params.BankType)
	}

	if len(params.Status) > 0 {
		dbModel = dbModel.Where("status in ?", params.Status)
	}

	var total int64
	if err := dbModel.Count(&total).Error; err != nil {
		return nil, err
	}

	var list []schema.WithdrawAccountUser

	err := dbModel.Preload("User").Limit(params.PageSize).Offset(utils.GetOffset(params.PageSize, params.Page)).Order("id desc").Find(&list).Error
	if err != nil {
		return nil, err
	}

	return &schema.WithdrawAccountListResponse{total, params.Page, len(list), list}, nil
}

func (s *WithdrawAccountService) WithdrawAccountUpdate(ids []int64, data map[string]interface{}) error {

	if err := s.db.Model(&model.WithdrawAccount{}).Where("id in ?", ids).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}

func (s *WithdrawAccountService) WithdrawAccountDel(id int64) error {

	if err := s.db.Delete(&model.WithdrawAccount{}, id).Error; err != nil {
		return err
	}

	return nil
}

func NewWithdrawAccountService(db *gorm.DB) *WithdrawAccountService {
	return &WithdrawAccountService{
		db: db,
	}
}
