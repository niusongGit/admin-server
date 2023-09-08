package service

import (
	"admin-server/internal/consts"
	"admin-server/internal/model"
	"admin-server/internal/schema"
	"admin-server/pkg/utils"
	"errors"
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

type PermissionService struct {
	db       *gorm.DB
	enforcer *casbin.SyncedCachedEnforcer
}

func (s *PermissionService) GetRole(query interface{}, args ...interface{}) (*model.Role, error) {

	m := model.Role{}

	if err := s.db.Where(query, args...).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (s *PermissionService) GetRoleList(params schema.RoleListRequest) (*schema.RoleListResponse, error) {

	dbModel := s.db.Model(&model.Role{})

	if params.Name != "" {
		dbModel = dbModel.Where("name LIKE ?", "%"+params.Name+"%")
	}

	var total int64
	if err := dbModel.Count(&total).Error; err != nil {
		return nil, err
	}

	var list []*model.Role

	err := dbModel.Limit(params.PageSize).Offset(utils.GetOffset(params.PageSize, params.Page)).Order("id desc").Find(&list).Error
	if err != nil {
		return nil, err
	}

	return &schema.RoleListResponse{total, params.Page, len(list), list}, nil
}

func (s *PermissionService) RoleAdd(params schema.RoleAddRequest) error {
	m := model.Role{
		Name: params.Name,
	}

	if err := s.db.Create(&m).Error; err != nil {
		return err
	}

	return nil
}

func (s *PermissionService) RoleUpdate(params schema.RoleUpdateRequest) error {
	data := map[string]interface{}{
		"name": params.Name,
	}

	if err := s.db.Model(&model.Role{}).Where("id = ?", params.Id).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}

func (s *PermissionService) RoleDel(id int64) error {
	if err := s.db.Delete(&model.Role{}, id).Error; err != nil {
		return err
	}
	s.enforcer.DeleteRole(fmt.Sprintf(consts.SubjectRole, id))
	return nil
}

func (s *PermissionService) GetSysApi(query interface{}, args ...interface{}) (*model.SysApi, error) {

	m := model.SysApi{}

	if err := s.db.Where(query, args...).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (s *PermissionService) GetSysApiList(params schema.SysApiListRequest) (*schema.SysApiListResponse, error) {

	dbModel := s.db.Model(&model.SysApi{})

	if params.Description != "" {
		dbModel = dbModel.Where("description LIKE ?", "%"+params.Description+"%")
	}

	if params.ApiGroup != "" {
		dbModel = dbModel.Where("api_group LIKE ?", "%"+params.ApiGroup+"%")
	}

	var total int64
	if err := dbModel.Count(&total).Error; err != nil {
		return nil, err
	}

	var list []*model.SysApi

	err := dbModel.Limit(params.PageSize).Offset(utils.GetOffset(params.PageSize, params.Page)).Order("id desc").Find(&list).Error
	if err != nil {
		return nil, err
	}

	return &schema.SysApiListResponse{total, params.Page, len(list), list}, nil
}

func (s *PermissionService) SysApiAdd(params schema.SysApiAddRequest) error {

	m := model.SysApi{}
	s.db.Where("path = ? and method = ?", params.Path, params.Method).First(&m)

	if m.Id > 0 {
		return errors.New("已经存在请勿重复添加！！")
	}

	m = model.SysApi{
		Path:        params.Path,
		Description: params.Description,
		ApiGroup:    params.ApiGroup,
		Method:      params.Method,
	}

	if err := s.db.Create(&m).Error; err != nil {
		return err
	}

	return nil
}

func (s *PermissionService) SysApiUpdate(params schema.SysApiUpdateRequest) error {

	var oldA model.SysApi
	err := s.db.Where("id = ?", params.Id).First(&oldA).Error
	if oldA.Path != params.Path || oldA.Method != params.Method {
		if !errors.Is(s.db.Where("path = ? AND method = ?", params.Path, params.Method).First(&model.SysApi{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同api路径")
		}
	}
	if err != nil {
		return err
	} else {
		err = s.UpdateCasbinApi(oldA.Path, params.Path, oldA.Method, params.Method)
		if err != nil {
			return err
		} else {
			data := map[string]interface{}{
				"path":        params.Path,
				"description": params.Description,
				"api_group":   params.ApiGroup,
				"method":      params.Method,
			}
			err = s.db.Model(&oldA).Where("id = ?", params.Id).Updates(&data).Error
		}
	}
	return err

}

func (s *PermissionService) SysApiDel(ids []int64) error {
	var apis []model.SysApi
	err := s.db.Find(&apis, "id in ?", ids).Delete(&apis).Error
	if err != nil {
		return err
	} else {
		for _, sysApi := range apis {
			s.ClearCasbin(1, sysApi.Path, sysApi.Method)
		}
		if err != nil {
			return err
		}
	}
	return err
}

func (s *PermissionService) UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	err := s.db.Model(&gormadapter.CasbinRule{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	err = s.enforcer.LoadPolicy()
	if err != nil {
		return err
	}
	return err
}

func (s *PermissionService) ClearCasbin(v int, p ...string) bool {
	success, _ := s.enforcer.RemoveFilteredPolicy(v, p...)
	return success
}

// 根据角色id获取角色的策略路径
func (s *PermissionService) GetPolicyPathByRoleId(roleId int64) (pathMaps []*schema.CasbinInfo) {
	list := s.enforcer.GetFilteredPolicy(0, fmt.Sprintf(consts.SubjectRole, roleId))
	for _, v := range list {
		pathMaps = append(pathMaps, &schema.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}

// 为角色设置权限
func (s *PermissionService) UpdateCasbinByRoleId(roleId int64, casbinInfos []*schema.CasbinInfo) error {
	authorityId := fmt.Sprintf(consts.SubjectRole, roleId)
	s.ClearCasbin(0, authorityId)
	rules := [][]string{}
	for _, v := range casbinInfos {
		rules = append(rules, []string{authorityId, v.Path, v.Method})
	}

	success, _ := s.enforcer.AddPolicies(rules)
	if !success {
		return errors.New("存在相同api,添加失败,请联系管理员")
	}
	return nil
}

func NewPermissionService(db *gorm.DB, enforcer *casbin.SyncedCachedEnforcer) *PermissionService {
	return &PermissionService{
		db:       db,
		enforcer: enforcer,
	}
}
