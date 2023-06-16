package logger

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"strings"
)

// 定义自己的Writer

type MyGormWriter struct {
	logger *zap.Logger
}

// 实现gorm/logger.Writer接口

func (m *MyGormWriter) Printf(format string, v ...interface{}) {
	fileLineStr, ok := v[0].(string)
	if ok {
		var tmpArr []string
		split := strings.Split(fileLineStr, "/")
		for i, str := range split {
			if i < len(split)-4 {
				continue
			}
			tmpArr = append(tmpArr, str)
		}
		v[0] = strings.Join(tmpArr, "/")
	}

	logstr := fmt.Sprintf(format, v...)
	m.logger.With(zap.String("log_from", "mysql")).Info(logstr)
}

func NewMyGormWriter() *MyGormWriter {
	return &MyGormWriter{logger: myLogger.Logger}
}

func NewMyGormWriterWithContext(c context.Context) *MyGormWriter {
	logger := GetLoggerByCtx(c)
	return &MyGormWriter{logger: logger}
}
