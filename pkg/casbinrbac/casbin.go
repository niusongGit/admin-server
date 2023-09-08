package casbinrbac

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

var (
	syncedCachedEnforcer *casbin.SyncedCachedEnforcer
)

// 初始化 Casbin Enforcer
func InitializeCasbinEnforcer(db *gorm.DB) (*casbin.SyncedCachedEnforcer, error) {

	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		return nil, err
	}
	//r.sub == "role_-1" 超级管理员不判断权限
	text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act || r.sub == "role_-1"
		`
	m, err := model.NewModelFromString(text)
	if err != nil {
		return nil, fmt.Errorf("字符串加载模型失败!", err.Error())
	}

	syncedCachedEnforcer, err = casbin.NewSyncedCachedEnforcer(m, adapter)
	if err != nil {
		return nil, err
	}
	syncedCachedEnforcer.SetExpireTime(60 * 60)
	_ = syncedCachedEnforcer.LoadPolicy()
	return syncedCachedEnforcer, nil
}

func GetCasbinEnforcer() *casbin.SyncedCachedEnforcer {
	return syncedCachedEnforcer
}
