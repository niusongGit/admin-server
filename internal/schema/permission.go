package schema

import "admin-server/internal/model"

type RoleIdRequest struct {
	Id int64 `json:"id" validate:"required|gt:0"`
}

type RoleListRequest struct {
	Name     string `json:"name" validate:"string|min_len:1|max_len:40"`
	Page     int    `json:"page"  validate:"gte:-1"`
	PageSize int    `json:"page_size" validate:"gte:-1"`
}

type RoleListResponse struct {
	Total    int64         `json:"total"`
	Page     int           `json:"page"`
	PageSize int           `json:"page_size"`
	List     []*model.Role `json:"list"`
}

type RoleAddRequest struct {
	Name string `json:"name" validate:"string|max_len:40"`
}

type RoleUpdateRequest struct {
	Id   int64  `json:"id" validate:"required|gt:0"`
	Name string `json:"name" validate:"string|max_len:40"`
}

type SysApiIdRequest struct {
	Id int64 `json:"id" validate:"required|gt:0"`
}

type SysApiListRequest struct {
	Description string `json:"description" validate:"string|min_len:1|max_len:40"`
	ApiGroup    string `json:"api_group" validate:"string|max_len:40"`
	Page        int    `json:"page"  validate:"gte:-1"`
	PageSize    int    `json:"page_size" validate:"gte:-1"`
}

type SysApiListResponse struct {
	Total    int64           `json:"total"`
	Page     int             `json:"page"`
	PageSize int             `json:"page_size"`
	List     []*model.SysApi `json:"list"`
}

type SysApiAddRequest struct {
	Description string `json:"description" validate:"required|string|min_len:1|max_len:40"`
	Path        string `json:"path" validate:"required|string|min_len:1"`
	ApiGroup    string `json:"api_group" validate:"required|string|max_len:40"`
	Method      string `json:"method" validate:"required|string|in:POST,GET,PUT,DELETE"` // 方法:创建POST(默认)|查看GET|更新PUT|删除DELETE
}

type SysApiUpdateRequest struct {
	Id          int64  `json:"id" validate:"required|gt:0"`
	Description string `json:"description" validate:"required|string|min_len:1|max_len:40"`
	Path        string `json:"path" validate:"required|string|min_len:1"`
	ApiGroup    string `json:"api_group" validate:"required|string|max_len:40"`
	Method      string `json:"method" validate:"required|string|in:POST,GET,PUT,DELETE"` // 方法:创建POST(默认)|查看GET|更新PUT|删除DELETE
}

type SysApiIdsRequest struct {
	Ids []int64 `json:"ids" validate:"required|ints"`
}

type SetRolePolicyPathRequest struct {
	Id              int64         `json:"id" validate:"required|gt:0"`
	RolePolicyPaths []*CasbinInfo `json:"role_policy_paths" alidate:"slice"`
}

// Casbin info structure
type CasbinInfo struct {
	Path   string `json:"path" validate:"required"`                                 // 路径
	Method string `json:"method" validate:"required|string|in:POST,GET,PUT,DELETE"` // 方法
}
