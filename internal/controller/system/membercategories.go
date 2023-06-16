package system

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

type MemberCategories struct {
}

func (c MemberCategories) Info(ctx *gin.Context) {

	data, err := service.NewSystemService(orm.GetDBWithContext(ctx.Request.Context())).GetMemberCategories()

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.FAILED_TO_QUERY_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, data, "成功")
}

func (c MemberCategories) Update(ctx *gin.Context) {

	var req = schema.MemberCategoriesRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		response.Fail(ctx, errmsg.INVALIDDATAFORMAT, nil)
		return
	}

	if err := validator.Validate(&req); err != nil {
		response.Response(ctx, http.StatusOK, errmsg.VALIDATEERROR, nil, err.Error())
		return
	}

	err := service.NewSystemService(orm.GetDBWithContext(ctx.Request.Context())).MemberCategoriesUpdate(req)

	if err != nil {
		response.Response(ctx, http.StatusOK, errmsg.ERROE_ADD_DATA, nil, err.Error())
		return
	}

	//返回结果
	response.Success(ctx, nil, "成功")
}

func NewMemberCategories() MemberCategories {
	return MemberCategories{}
}
