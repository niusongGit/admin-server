package service

import (
	"admin-server/internal/model"
	"admin-server/internal/schema"
	"admin-server/pkg/utils"
	"gorm.io/gorm"
)

type AnnouncementService struct {
	db *gorm.DB
}

func (s *AnnouncementService) GetAnnouncement(query interface{}, args ...interface{}) (*schema.AnnouncementSportType, error) {

	m := schema.AnnouncementSportType{}

	if err := s.db.Where(query, args...).Preload("SportType").First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (s *AnnouncementService) GetAnnouncementList(params schema.AnnouncementListRequest) (*schema.AnnouncementListResponse, error) {

	dbModel := s.db.Model(&model.Announcement{})

	if params.Title != "" {
		dbModel = dbModel.Where("title LIKE ?", "%"+params.Title+"%")
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

	var list []*schema.AnnouncementSportType

	err := dbModel.Preload("SportType").Limit(params.PageSize).Offset(utils.GetOffset(params.PageSize, params.Page)).Order("id desc").Find(&list).Error
	if err != nil {
		return nil, err
	}

	return &schema.AnnouncementListResponse{total, params.Page, len(list), list}, nil
}

func (s *AnnouncementService) AnnouncementAdd(params schema.AnnouncementAddRequest) error {
	m := model.Announcement{
		SportTypeId: params.SportTypeId,
		Title:       params.Title,
		Content:     params.Content,
		Status:      params.Status,
	}

	if err := s.db.Create(&m).Error; err != nil {
		return err
	}

	return nil
}

func (s *AnnouncementService) AnnouncementUpdate(params schema.AnnouncementUpdateRequest) error {
	data := map[string]interface{}{
		"sport_type_id": params.SportTypeId,
		"title":         params.Title,
		"content":       params.Content,
		"status":        params.Status,
	}

	if err := s.db.Model(&model.Announcement{}).Where("id = ?", params.Id).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}

func (s *AnnouncementService) AnnouncementDel(id int64) error {

	if err := s.db.Delete(&model.Announcement{}, id).Error; err != nil {
		return err
	}

	return nil
}

func NewAnnouncementService(db *gorm.DB) *AnnouncementService {
	return &AnnouncementService{
		db: db,
	}
}
