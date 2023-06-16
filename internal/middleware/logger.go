package middleware

import (
	"admin-server/pkg/logger"
	"admin-server/pkg/utils"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"strings"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

// Gin中间件函数，记录请求日志

func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter
		// 开始时间
		startTime := time.Now()

		bodyData, _ := c.GetRawData()

		// 重新赋值
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyData))

		// 向日志中添加traceId，在每个请求中加入追踪编号
		traceId := fmt.Sprintf("%s-%v", utils.RandomString(20), startTime.UnixMilli())

		log, ctx := logger.AddCtxWithTraceId(c, traceId)
		c.Request = c.Request.WithContext(ctx)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 请求IP
		clientIP := c.ClientIP()

		//日志格式
		bodyStr := strings.Replace(string(bodyData), "\n", "", -1)
		log = log.With(
			zap.String("log_from", "http_middleware"),
			zap.String("uri", reqUri),
			zap.String("ip", clientIP),
			zap.String("method", reqMethod),
			zap.String("request_data", bodyStr),
		)
		log.Info("")
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()

		//执行时间
		latencyTime := fmt.Sprintf("%6v", endTime.Sub(startTime))

		// 状态码
		statusCode := c.Writer.Status()

		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}

		log = log.With(
			zap.String("log_from", "http_middleware"),
			zap.String("uri", reqUri),
			zap.Int("http_status", statusCode),
			zap.String("total_time", latencyTime),
			zap.String("ip", clientIP),
			zap.String("method", reqMethod),
			zap.Int("data_size", dataSize),
			zap.String("response_data", bodyLogWriter.body.String()),
		)

		if len(c.Errors) > 0 {
			log.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			log.Error("")
		} else if statusCode >= 400 {
			log.Warn("")
		} else {
			log.Info("")
		}

	}
}
