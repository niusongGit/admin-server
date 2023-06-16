package service

import (
	"admin-server/internal/consts"
	"admin-server/internal/model"
	"admin-server/internal/schema"
	"admin-server/pkg/utils"
	"gorm.io/gorm"
)

type ExpertApplicationAuditService struct {
	db *gorm.DB
}

func (s *ExpertApplicationAuditService) GetApplicationAuditList(params schema.ApplicationAuditListRequest) (*schema.ApplicationAuditListResponse, error) {

	dbModel := s.db.Model(&model.ExpertApplicationAudit{}).
		Preload("User", func(db *gorm.DB) *gorm.DB { return db.Model(&model.User{}) }).
		Preload("SportType", func(db *gorm.DB) *gorm.DB { return db.Model(model.SportType{}) })
	if len(params.Statuses) > 0 {
		dbModel.Where("status in (?)", params.Statuses)
	}

	var total int64
	if err := dbModel.Count(&total).Error; err != nil {
		return nil, err
	}

	var list []*schema.UserExpertApplicationAudit

	err := dbModel.Limit(params.PageSize).Offset(utils.GetOffset(params.PageSize, params.Page)).
		Order("id desc").Find(&list).Error
	if err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return &schema.ApplicationAuditListResponse{total, params.Page, len(list), list}, nil
	}

	var ids []int64
	for _, v := range list {
		ids = append(ids, v.Id)
	}

	return &schema.ApplicationAuditListResponse{total, params.Page, len(list), list}, nil
}

func (s *ExpertApplicationAuditService) ApplicationAuditUpdate(params schema.ApplicationAuditUpdateRequest) error {

	if err := s.db.Model(&model.User{}).Where("id = ?", params.Id).Updates(&model.ExpertApplicationAudit{
		Id: params.Id,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (s *ExpertApplicationAuditService) ApplicationAuditApprove(params schema.ApplicationAuditApproveRequest) error {
	eaa := new(model.ExpertApplicationAudit)
	if err := s.db.First(eaa, "id = ?", params.Id).Error; err != nil {
		return err
	}

	status := consts.StatusFailed
	isExpert := 0
	if params.IsApproved {
		status = consts.StatusSuccess
		isExpert = 1
	}

	tx := s.db.Begin()
	defer tx.Rollback()

	err := tx.Model(&model.ExpertApplicationAudit{}).Where("id = ?", params.Id).
		Updates(&model.ExpertApplicationAudit{
			Id:     params.Id,
			Status: int64(status),
			Remark: params.Remark,
		}).Error
	if err != nil {
		return err
	}

	err = tx.Model(&model.User{}).Where("id = ?", eaa.UserId).Select("is_expert", "sport_type_id").
		Updates(&model.User{SportTypeId: eaa.SportTypeId, IsExpert: int64(isExpert)}).Error
	if err != nil {
		return err
	}
	tx.Commit()

	return nil
}

func NewExpertApplicationAuditService(db *gorm.DB) *ExpertApplicationAuditService {
	return &ExpertApplicationAuditService{
		db: db,
	}
}
