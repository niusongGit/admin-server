package service

import (
	"admin-server/internal/consts"
	"admin-server/internal/model"
	"admin-server/internal/schema"
	"admin-server/pkg/orm/datatypes"
	"admin-server/pkg/utils"
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

func (s *AdminService) GetAdminList(params schema.AdminListRequest) (*schema.AdminListResponse, error) {

	dbModel := s.db.Model(&model.Admin{})

	if params.AdminName != "" {
		dbModel = dbModel.Where("admin_name LIKE ?", "%"+params.AdminName+"%")
	}

	if len(params.RoleIds) > 0 {
		dbModel = dbModel.Where("role_id in ?", params.RoleIds)
	}

	var total int64
	if err := dbModel.Count(&total).Error; err != nil {
		return nil, err
	}

	var list []*schema.AdminList

	err := dbModel.Preload("Role").Limit(params.PageSize).Offset(utils.GetOffset(params.PageSize, params.Page)).Order("id desc").Find(&list).Error
	if err != nil {
		return nil, err
	}

	for i, v := range list {
		if v.RoleId == -1 {
			list[i].Role.Id = -1
			list[i].Role.Name = consts.RootRoleName
		}
	}

	return &schema.AdminListResponse{total, params.Page, len(list), list}, nil
}

func (s *AdminService) AdminUpdate(params schema.AdminRequest) error {
	data := map[string]interface{}{
		"role_id":    params.RoleId,
		"admin_name": params.AdminName,
		"password":   params.Password,
	}

	if err := s.db.Model(&model.Admin{}).Where("id = ?", params.Id).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}

func NewAdminService(db *gorm.DB) *AdminService {
	return &AdminService{
		db: db,
	}
}
