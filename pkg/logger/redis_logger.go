package logger

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"strings"
)

// 定义自己的Writer

type MyRedisLogger struct {
	logger *zap.Logger
	count  int
}

func (m *MyRedisLogger) Printf(ctx context.Context, format string, v ...interface{}) {
	logstr := fmt.Sprintf(format, v...)
	if strings.Contains(logstr, "redis: discarding bad PubSub connection: ") || strings.Contains(logstr, ": i/o timeout") {
		m.count++
	}
	if m.count > 3 {
		return
	}
	m.logger.With(zap.String("log_from", "redis")).Info(logstr)
}

// 实现gorm/logger.Writer接口

func NewMyRedisLogger() *MyRedisLogger {
	return &MyRedisLogger{logger: myLogger.Logger}
}

func NewMyRedisLoggerWithContext(c context.Context) *MyRedisLogger {
	logger := GetLoggerByCtx(c)
	return &MyRedisLogger{logger: logger}
}
