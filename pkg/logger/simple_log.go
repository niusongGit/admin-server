package logger

import "go.uber.org/zap"

var Sugar *zap.SugaredLogger

func Debug(msg string, fields ...zap.Field) {
	myLogger.Logger.WithOptions(zap.AddCallerSkip(1)).Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	myLogger.Logger.WithOptions(zap.AddCallerSkip(1)).Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	myLogger.Logger.WithOptions(zap.AddCallerSkip(1)).Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	myLogger.Logger.WithOptions(zap.AddCallerSkip(1)).Error(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	myLogger.Logger.WithOptions(zap.AddCallerSkip(1)).DPanic(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	myLogger.Logger.WithOptions(zap.AddCallerSkip(1)).Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	myLogger.Logger.WithOptions(zap.AddCallerSkip(1)).Fatal(msg, fields...)
}

func With(fields ...zap.Field) *zap.Logger {
	return myLogger.Logger.With(fields...)
}
