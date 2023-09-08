package permission

import (
	"admin-server/internal/response"
	"admin-server/internal/schema"
	"admin-server/internal/service"
	"admin-server/pkg/casbinrbac"
	"admin-server/pkg/errmsg"
	"admin-server/pkg/orm"
	"admin-server/pkg/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Role struct {
}

func (c Role) Info(ctx *gin.Context) {

	var req = schema.RoleIdRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	data, err := service.NewPermissionService(orm.GetDBWithContext(ctx.Request.Context()), casbinrbac.GetCasbinEnforcer()).GetRole("id = ?", req.Id)

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.FAILED_TO_QUERY_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, data, "成功")
}

func (c Role) List(ctx *gin.Context) {

	var req = schema.RoleListRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	data, err := service.NewPermissionService(orm.GetDBWithContext(ctx.Request.Context()), casbinrbac.GetCasbinEnforcer()).GetRoleList(req)

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.FAILED_TO_QUERY_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, data, "成功")
}

func (c Role) Add(ctx *gin.Context) {

	var req = schema.RoleAddRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	err := service.NewPermissionService(orm.GetDBWithContext(ctx.Request.Context()), casbinrbac.GetCasbinEnforcer()).RoleAdd(req)

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.ERROE_ADD_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, nil, "成功")
}

func (c Role) Update(ctx *gin.Context) {

	var req = schema.RoleUpdateRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	err := service.NewPermissionService(orm.GetDBWithContext(ctx.Request.Context()), casbinrbac.GetCasbinEnforcer()).RoleUpdate(req)

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.ERROE_ADD_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, nil, "成功")
}

func (c Role) Del(ctx *gin.Context) {

	var req = schema.RoleIdRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	err := service.NewPermissionService(orm.GetDBWithContext(ctx.Request.Context()), casbinrbac.GetCasbinEnforcer()).RoleDel(req.Id)

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.ERROE_DELETE_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, nil, "成功")
}

// 获取一个角色的权限

func (c Role) GetRolePolicyPath(ctx *gin.Context) {

	var req = schema.RoleIdRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	list := service.NewPermissionService(orm.GetDBWithContext(ctx.Request.Context()), casbinrbac.GetCasbinEnforcer()).GetPolicyPathByRoleId(req.Id)

	//返回结果
	response.Success(ctx, list, "成功")
}

// 设置一个角色的权限

func (c Role) SetRolePolicyPath(ctx *gin.Context) {

	var req = schema.SetRolePolicyPathRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	err := service.NewPermissionService(orm.GetDBWithContext(ctx.Request.Context()), casbinrbac.GetCasbinEnforcer()).UpdateCasbinByRoleId(req.Id, req.RolePolicyPaths)
	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.ERROE_DELETE_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, nil, "成功")
}

func NewRole() Role {
	return Role{}
}
