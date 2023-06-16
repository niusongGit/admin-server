package payment

import (
	"admin-server/internal/response"
	"admin-server/internal/schema"
	"admin-server/internal/service"
	"admin-server/pkg/errmsg"
	"admin-server/pkg/orm"
	"admin-server/pkg/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WithdrawApplication struct {
}

func (c WithdrawApplication) Info(ctx *gin.Context) {

	var req = schema.WithdrawApplicationIdRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	data, err := service.NewWithdrawApplicationService(orm.GetDBWithContext(ctx.Request.Context())).GetWithdrawApplication("id = ?", req.Id)

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.FAILED_TO_QUERY_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, data, "成功")
}

func (c WithdrawApplication) List(ctx *gin.Context) {

	var req = schema.WithdrawApplicationListRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	data, err := service.NewWithdrawApplicationService(orm.GetDBWithContext(ctx.Request.Context())).GetWithdrawApplicationList(req)

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.FAILED_TO_QUERY_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, data, "成功")
}

func (c WithdrawApplication) Audit(ctx *gin.Context) {

	var req = schema.WithdrawApplicationAuditRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	err := service.NewWithdrawApplicationService(orm.GetDBWithContext(ctx.Request.Context())).WithdrawApplicationAudit(req)

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.ERROE_ADD_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, nil, "成功")
}

func (c WithdrawApplication) Del(ctx *gin.Context) {

	var req = schema.WithdrawApplicationIdRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	err := service.NewWithdrawApplicationService(orm.GetDBWithContext(ctx.Request.Context())).WithdrawApplicationDel(req.Id)

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.ERROE_DELETE_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, nil, "成功")
}

func NewWithdrawApplication() WithdrawApplication {
	return WithdrawApplication{}
}
