package middleware

import (
	"admin-server/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"runtime/debug"
)

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			//log.Printf("panic: %v\n", r)
			debug.PrintStack()

			// 请求IP
			clientIP := c.ClientIP()
			//日志格式
			logger.GetLoggerByCtx(c.Request.Context()).WithOptions(zap.AddCallerSkip(2)).With(
				zap.String("log_from", "recover_middleware"),
				zap.String("uri", c.Request.RequestURI),
				zap.Int("http_status", c.Writer.Status()),
				zap.String("ip", clientIP),
				zap.String("method", c.Request.Method),
			).Error("", zap.Any("err", r))

			c.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  errorToString(r),
				"data": nil,
			})
			//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}

// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
