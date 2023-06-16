package service

import (
	"admin-server/internal/model"
	"admin-server/internal/schema"
	"admin-server/pkg/utils"
	"gorm.io/gorm"
)

type VersionService struct {
	db *gorm.DB
}

func (s *VersionService) GetVersion(query interface{}, args ...interface{}) (*model.Version, error) {

	m := model.Version{}

	if err := s.db.Where(query, args...).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (s *VersionService) GetVersionList(params schema.VersionListRequest) (*schema.VersionListResponse, error) {

	dbModel := s.db.Model(&model.Version{})

	if params.Version != "" {
		dbModel = dbModel.Where("version = ?", params.Version)
	}

	if params.VersionNumber > 0 {
		dbModel = dbModel.Where("version_number = ?", params.VersionNumber)
	}
	if len(params.Force) > 0 {
		dbModel = dbModel.Where("force in ?", params.Force)
	}
	if len(params.Status) > 0 {
		dbModel = dbModel.Where("status in ?", params.Status)
	}

	var total int64
	if err := dbModel.Count(&total).Error; err != nil {
		return nil, err
	}

	var list []model.Version

	err := dbModel.Limit(params.PageSize).Offset(utils.GetOffset(params.PageSize, params.Page)).Order("id desc").Find(&list).Error
	if err != nil {
		return nil, err
	}

	return &schema.VersionListResponse{total, params.Page, len(list), list}, nil
}

func (s *VersionService) VersionAdd(params schema.VersionAddRequest) error {
	m := model.Version{
		Version:       params.Version,
		VersionNumber: params.VersionNumber,
		UpdateUrl:     params.UpdateUrl,
		UpdateBin:     params.UpdateBin,
		UpdateLog:     params.UpdateLog,
		Force:         params.Force,
		Status:        params.Status,
	}

	if err := s.db.Create(&m).Error; err != nil {
		return err
	}

	return nil
}

func (s *VersionService) VersionUpdate(params schema.VersionUpdateRequest) error {

	data := map[string]interface{}{
		"version":        params.Version,
		"version_number": params.VersionNumber,
		"update_url":     params.UpdateUrl,
		"update_bin":     params.UpdateBin,
		"update_log":     params.UpdateLog,
		"force":          params.Force,
		"status":         params.Status,
	}

	if err := s.db.Model(&model.Version{}).Where("id = ?", params.Id).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}

func (s *VersionService) VersionDel(id int64) error {

	if err := s.db.Delete(&model.Version{}, id).Error; err != nil {
		return err
	}

	return nil
}

func NewVersionService(db *gorm.DB) *VersionService {
	return &VersionService{
		db: db,
	}
}
