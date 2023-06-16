package service

import (
	"admin-server/internal/consts"
	"admin-server/internal/model"
	"admin-server/internal/schema"
	"admin-server/pkg/utils"
	"errors"
	"gorm.io/gorm"
)

type WithdrawApplicationService struct {
	db *gorm.DB
}

func (s *WithdrawApplicationService) GetWithdrawApplication(query interface{}, args ...interface{}) (*schema.WithdrawApplicationUser, error) {

	m := schema.WithdrawApplicationUser{}

	if err := s.db.Model(model.WithdrawApplication{}).Preload("User").Preload("WithdrawAccount").Where(query, args...).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (s *WithdrawApplicationService) GetWithdrawApplicationList(params schema.WithdrawApplicationListRequest) (*schema.WithdrawApplicationListResponse, error) {

	dbModel := s.db.Model(&model.WithdrawApplication{})

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

	withdrawAccountWhere := make(map[string]interface{})

	if params.Name != "" {
		withdrawAccountWhere["name"] = params.Name
	}

	if params.Number != "" {
		withdrawAccountWhere["number"] = params.Number
	}

	if params.Type != "" {
		withdrawAccountWhere["type"] = params.Type
	}

	if params.BankType != "" {
		withdrawAccountWhere["bank_type"] = params.BankType
	}

	if len(withdrawAccountWhere) > 0 {
		withdrawAccountIds := make([]int64, 0)
		s.db.Model(&model.WithdrawAccount{}).Where(withdrawAccountWhere).Pluck("id", &withdrawAccountIds)
		dbModel = dbModel.Where("withdraw_account_id in ?", withdrawAccountIds)
	}

	if len(params.Status) > 0 {
		dbModel = dbModel.Where("status in ?", params.Status)
	}

	var total int64
	if err := dbModel.Count(&total).Error; err != nil {
		return nil, err
	}

	var list []schema.WithdrawApplicationUser

	err := dbModel.Preload("User").Preload("WithdrawAccount").Limit(params.PageSize).Offset(utils.GetOffset(params.PageSize, params.Page)).Order("id desc").Find(&list).Error
	if err != nil {
		return nil, err
	}

	return &schema.WithdrawApplicationListResponse{total, params.Page, len(list), list}, nil
}

func (s *WithdrawApplicationService) WithdrawApplicationAudit(params schema.WithdrawApplicationAuditRequest) error {
	return s.db.Transaction(func(tx *gorm.DB) error {

		var withdrawApplication model.WithdrawApplication

		if err := tx.First(&withdrawApplication, params.Id).Error; err != nil {
			return errors.New("查询提现申请记录失败" + err.Error())
		}

		if (consts.WithdrawStatusFailed == params.Status || consts.WithdrawStatusSuccess == params.Status) && consts.WithdrawStatusWaiting != withdrawApplication.Status {
			return errors.New("只有待审核状态下才能进行审核操作")
		}

		if err := tx.Model(&model.WithdrawApplication{}).Where("id = ?", params.Id).Updates(map[string]interface{}{
			"status": params.Status,
			"remark": params.Remark,
		}).Error; err != nil {
			return errors.New("审核提现申请失败：" + err.Error())
		}

		if consts.WithdrawStatusFailed == params.Status {
			if err := tx.Model(&model.User{Id: withdrawApplication.UserId}).Update("withdraw_amount", gorm.Expr("withdraw_amount + ?", withdrawApplication.WithdrawAmount)).Error; err != nil {
				return errors.New("审核提现申请未通过，返还用户提现金额失败：" + err.Error())
			}
		}

		return nil
	})
}

func (s *WithdrawApplicationService) WithdrawApplicationDel(id int64) error {

	if err := s.db.Delete(&model.WithdrawApplication{}, id).Error; err != nil {
		return err
	}

	return nil
}

func NewWithdrawApplicationService(db *gorm.DB) *WithdrawApplicationService {
	return &WithdrawApplicationService{
		db: db,
	}
}
