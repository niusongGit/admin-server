package logger

import (
	"context"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)

var myLogger *MyLogger

const ctxKey = "trace_id_log"
const traceIdKey = "trace_id"

type MyLogger struct {
	*zap.Logger
}

func GetLoggerByCtx(ctx context.Context) *zap.Logger {
	log, ok := ctx.Value(ctxKey).(*zap.Logger)
	if !ok {
		return myLogger.Logger
	}
	return log
}

func GetTraceIdByCtx(ctx context.Context) string {
	traceId, ok := ctx.Value(traceIdKey).(string)
	if !ok {
		return ""
	}
	return traceId
}

func AddCtxWithTraceId(ctx context.Context, traceId string) (logger *zap.Logger, ctx1 context.Context) {
	logger = myLogger.With(zap.String(traceIdKey, traceId))
	ctx1 = context.WithValue(ctx, traceIdKey, traceId)
	ctx1 = context.WithValue(ctx1, ctxKey, logger)
	return
}

func AddCtx(ctx context.Context, field ...zap.Field) (logger *zap.Logger, ctx1 context.Context) {
	logger = myLogger.With(field...)
	ctx1 = context.WithValue(ctx, ctxKey, logger)
	return
}

func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段,如：添加一个服务器名称
	filed := zap.Fields(zap.String("serviceName", "admin_service"))
	// 构造日志
	logger := zap.New(core, caller, development, filed)
	logger.Info("DefaultLogger init success")
	myLogger = &MyLogger{logger}
	Sugar = myLogger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:    "msg",
		LevelKey:      "level",
		TimeKey:       "time",
		NameKey:       "logger",
		CallerKey:     "line",
		FunctionKey:   zapcore.OmitKey,
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalLevelEncoder, // 大写编码器
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			layout := "2006-01-02 15:04:05"
			type appendTimeEncoder interface {
				AppendTimeLayout(time.Time, string)
			}
			if enc, ok := enc.(appendTimeEncoder); ok {
				enc.AppendTimeLayout(t, layout)
				return
			}
			enc.AppendString(t.Format(layout))
		},
		//EncodeDuration: zapcore.StringDurationEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	//encoderConfig = zap.NewDevelopmentEncoderConfig()
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	path := "./logs/log.log"
	//path := "../../logs/log.log"
	l, _ := rotatelogs.New(
		path+"_%Y%m%d",
		rotatelogs.WithLinkName(path),             // 生成软链，指向最新日志文件
		rotatelogs.WithRotationCount(15),          // 文件最多存在的数量
		rotatelogs.WithRotationTime(time.Hour*24), // 24小时切割一次
	)
	// 利用io.MultiWriter支持文件和终端两个输出目标
	ws := io.MultiWriter(l, os.Stdout)
	return zapcore.AddSync(ws)

}
