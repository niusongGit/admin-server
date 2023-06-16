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

type WithdrawAccount struct {
}

func (c WithdrawAccount) Info(ctx *gin.Context) {

	var req = schema.WithdrawAccountIdRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	data, err := service.NewWithdrawAccountService(orm.GetDBWithContext(ctx.Request.Context())).GetWithdrawAccount("id = ?", req.Id)

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.FAILED_TO_QUERY_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, data, "成功")
}

func (c WithdrawAccount) List(ctx *gin.Context) {

	var req = schema.WithdrawAccountListRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	data, err := service.NewWithdrawAccountService(orm.GetDBWithContext(ctx.Request.Context())).GetWithdrawAccountList(req)

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.FAILED_TO_QUERY_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, data, "成功")
}

func (c WithdrawAccount) Update(ctx *gin.Context) {

	var req = schema.WithdrawAccountUpdateRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	err := service.NewWithdrawAccountService(orm.GetDBWithContext(ctx.Request.Context())).WithdrawAccountUpdate([]int64{req.Id}, map[string]interface{}{
		"name":      req.Name,
		"number":    req.Number,
		"type":      req.Type,
		"bank_type": req.BankType,
		"status":    req.Status,
	})

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.ERROE_ADD_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, nil, "成功")
}

func (c WithdrawAccount) Del(ctx *gin.Context) {

	var req = schema.WithdrawAccountIdRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	err := service.NewWithdrawAccountService(orm.GetDBWithContext(ctx.Request.Context())).WithdrawAccountDel(req.Id)

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.ERROE_DELETE_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, nil, "成功")
}

func NewWithdrawAccount() WithdrawAccount {
	return WithdrawAccount{}
}
