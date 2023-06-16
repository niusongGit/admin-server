package response

import (
	"admin-server/pkg/errmsg"
	"admin-server/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 统一返回数据json格式
func Response(ctx *gin.Context, httpStatus int, code int, data interface{}, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg, "request_id": logger.GetTraceIdByCtx(ctx.Request.Context())})
}

//成功

func Success(ctx *gin.Context, data interface{}, msg string) {
	Response(ctx, http.StatusOK, errmsg.SUCCSE, data, msg)
}

//失败

func Fail(ctx *gin.Context, code int, data interface{}) {
	Response(ctx, http.StatusOK, code, data, errmsg.GetErrMsg(code))
}
