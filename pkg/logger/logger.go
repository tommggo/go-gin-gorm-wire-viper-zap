package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"go-gin-gorm-wire-viper-zap/internal/config"
)

// Logger 定义了通用的日志接口
type Logger interface {
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)
	With(fields ...Field) Logger
	Printf(format string, args ...interface{})
}

// Field 定义日志字段
type Field = zap.Field

// zapLogger 实现 Logger 接口
type zapLogger struct {
	l *zap.Logger
}

// 确保 zapLogger 实现了 Logger 接口
var _ Logger = (*zapLogger)(nil)

func (l *zapLogger) Debug(msg string, fields ...Field) { l.l.Debug(msg, fields...) }
func (l *zapLogger) Info(msg string, fields ...Field)  { l.l.Info(msg, fields...) }
func (l *zapLogger) Warn(msg string, fields ...Field)  { l.l.Warn(msg, fields...) }
func (l *zapLogger) Error(msg string, fields ...Field) { l.l.Error(msg, fields...) }
func (l *zapLogger) Fatal(msg string, fields ...Field) { l.l.Fatal(msg, fields...) }
func (l *zapLogger) With(fields ...Field) Logger       { return &zapLogger{l.l.With(fields...)} }
func (l *zapLogger) Printf(format string, args ...interface{}) {
	l.Info(fmt.Sprintf(format, args...))
}

// 全局日志实例
var std Logger

// 提供便捷的全局方法
func Debug(msg string, fields ...Field)         { std.Debug(msg, fields...) }
func Info(msg string, fields ...Field)          { std.Info(msg, fields...) }
func Warn(msg string, fields ...Field)          { std.Warn(msg, fields...) }
func Error(msg string, fields ...Field)         { std.Error(msg, fields...) }
func Fatal(msg string, fields ...Field)         { std.Fatal(msg, fields...) }
func With(fields ...Field) Logger               { return std.With(fields...) }
func Printf(format string, args ...interface{}) { std.Printf(format, args...) }

// 提供常用的字段构造方法
var (
	String   = zap.String
	Int      = zap.Int
	Int64    = zap.Int64
	Float64  = zap.Float64
	Bool     = zap.Bool
	Err      = zap.Error
	Any      = zap.Any
	Duration = zap.Duration
	Time     = zap.Time
)

// Setup 初始化日志
func Setup(cfg *config.Config) {
	// 1. 设置日志级别
	level, err := zapcore.ParseLevel(cfg.Log.Level)
	if err != nil {
		panic(fmt.Sprintf("parse log level failed: %v", err))
	}

	// 2. 配置日志输出
	fileWriter := &lumberjack.Logger{
		Filename:   cfg.Log.Filename,
		MaxSize:    cfg.Log.MaxSize,
		MaxBackups: cfg.Log.MaxBackups,
		MaxAge:     cfg.Log.MaxAge,
		Compress:   cfg.Log.Compress,
	}

	// 3. 配置编码器
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// 4. 创建核心
	core := zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.AddSync(fileWriter),
			level,
		),
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.AddSync(os.Stdout),
			level,
		),
	)

	// 5. 创建日志实例
	l := zap.New(core,
		zap.AddCaller(),
		zap.AddCallerSkip(2),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)

	// 6. 设置全局实例
	std = &zapLogger{l: l}
}

// 获取标准 logger
func StandardLogger() Logger {
	return std
}
