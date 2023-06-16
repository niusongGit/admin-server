package service

import (
	"admin-server/internal/model"
	"admin-server/pkg/orm/datatypes"
	"errors"
	"gorm.io/gorm"
	"time"
)

type AdminService struct {
	db *gorm.DB
}

func (s *AdminService) Add(m *model.Admin) error {
	if err := s.db.Create(m).Error; err != nil {
		return errors.New("添加管理员数据失败：" + err.Error())
	}
	return nil
}

func (s *AdminService) UpdateLastLoinTime(id int64) error {
	if err := s.db.Model(&model.Admin{}).Where("id = ?", id).UpdateColumn("last_login_time", datatypes.XTime{time.Now()}).Error; err != nil {
		return errors.New("保存管理员数据失败：" + err.Error())
	}
	return nil
}

func (s *AdminService) IsAdminExist(adminName string) bool {
	var admin model.Admin
	if res := s.db.Where("admin_name = ?", adminName).First(&admin); res.Error != nil {
		return false
	}
	return true
}

func (s *AdminService) ChangePassword(id int64, passWord string) error {
	if err := s.db.Model(&model.Admin{}).Where("id = ?", id).Update("password", passWord).Error; err != nil {
		return err
	}
	return nil
}

func NewAdminService(db *gorm.DB) *AdminService {
	return &AdminService{
		db: db,
	}
}
