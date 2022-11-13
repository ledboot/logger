package logger

import (
	"context"
	"fmt"
	"os"
	"time"

	gormlogger "gorm.io/gorm/logger"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
	Debugw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
	Log(level Level, args ...interface{})
	Logf(level Level, template string, args ...interface{})
	Logw(level Level, msg string, keysAndValues ...interface{})
}

type logger struct {
	config *Config
	logger *zap.SugaredLogger
}

var (
	defaultConfig = NewDefaultConfig()
	l, ZapLogger  = newLogger(defaultConfig)
	gormLogger    = NewGormLogger(context.Background(), l, 5*time.Second)
)

func SetConfig(config *Config) {
	l, ZapLogger = newLogger(config)
	gormLogger = NewGormLogger(context.Background(), l, 5*time.Second)
}

func SetLevel(level Level) {
	l.SetLevel(level)
}

func GetLevel() Level {
	return l.Level()
}

func SetOutputPaths(outputPaths []string) {
	l.config.OutputPaths = outputPaths
	l, ZapLogger = newLogger(l.config)
}

func With(key string, value interface{}) {
	l.config.InitialFields[key] = value
	l, ZapLogger = newLogger(l.config)
}

func NewLogger(config *Config) Logger {
	l, _ := newLogger(l.config)
	return l
}

func GetLogger() Logger {
	return l
}

func GetGormLogger() gormlogger.Interface {
	return gormLogger
}

func newLogger(config *Config) (*logger, *zap.Logger) {
	encoderConfig := NewCustomEncoderConfig(config)
	zapConfig := &zap.Config{
		Level:             zap.NewAtomicLevelAt(zapcore.Level(config.Level)),
		Development:       config.Development,
		DisableCaller:     config.DisableCaller,
		DisableStacktrace: config.DisableStacktrace,
		Sampling:          &zap.SamplingConfig{Initial: 100, Thereafter: 100},
		Encoding:          config.Encoding,
		EncoderConfig:     encoderConfig,
		OutputPaths:       config.OutputPaths,
		InitialFields:     config.InitialFields,
	}

	zapLogger, err := zapConfig.Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error on build zap logger (%s)", err)
		return nil, nil
	}
	zapLogger = zapLogger.WithOptions(zap.AddCallerSkip(config.CallerSkip))
	config.zapConfig = zapConfig

	return &logger{
		logger: zapLogger.Sugar(),
		config: config,
	}, zapLogger
}

func NewCustomEncoderConfig(conf *Config) zapcore.EncoderConfig {
	encodeLevel := zapcore.LowercaseLevelEncoder
	if conf.EnableColor {
		encodeLevel = zapcore.LowercaseColorLevelEncoder
	}
	encodeTime := zapcore.ISO8601TimeEncoder
	if conf.ShortTime {
		encodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			type appendTimeEncoder interface {
				AppendTimeLayout(time.Time, string)
			}
			layout := "2006-01-02 15:04:05"
			if enc, ok := enc.(appendTimeEncoder); ok {
				enc.AppendTimeLayout(t, layout)
				return
			}
			enc.AppendString(t.Format(layout))
		}
	}
	return zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    encodeLevel,
		EncodeTime:     encodeTime,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func (l *logger) SetLevel(level Level) {
	l.config.Level = level
	l.config.zapConfig.Level.SetLevel(zapcore.Level(level))
}

func (l *logger) Level() Level {
	return l.config.Level
}

func (l *logger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *logger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *logger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *logger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *logger) DPanic(args ...interface{}) {
	l.logger.DPanic(args...)
}

func (l *logger) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

func (l *logger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *logger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

func (l *logger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

func (l *logger) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}

func (l *logger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

func (l *logger) DPanicf(template string, args ...interface{}) {
	l.logger.DPanicf(template, args...)
}

func (l *logger) Panicf(template string, args ...interface{}) {
	l.logger.Panicf(template, args...)
}

func (l *logger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args...)
}

func (l *logger) Debugw(msg string, keysAndValues ...interface{}) {
	l.logger.Debugw(msg, keysAndValues...)
}

func (l *logger) Infow(msg string, keysAndValues ...interface{}) {
	l.logger.Infow(msg, keysAndValues...)
}

func (l *logger) Warnw(msg string, keysAndValues ...interface{}) {
	l.logger.Warnw(msg, keysAndValues...)
}

func (l *logger) Errorw(msg string, keysAndValues ...interface{}) {
	l.logger.Errorw(msg, keysAndValues...)
}

func (l *logger) DPanicw(msg string, keysAndValues ...interface{}) {
	l.logger.DPanicw(msg, keysAndValues...)
}

func (l *logger) Panicw(msg string, keysAndValues ...interface{}) {
	l.logger.Panicw(msg, keysAndValues...)
}

func (l *logger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.logger.Fatalw(msg, keysAndValues...)
}

func (l *logger) Log(level Level, args ...interface{}) {
	switch level {
	case DebugLevel:
		l.logger.Debug(args...)
	case InfoLevel:
		l.logger.Info(args...)
	case WarnLevel:
		l.logger.Warn(args...)
	case ErrorLevel:
		l.logger.Error(args...)
	default:
		l.logger.Info(args...)
	}
}

func (l *logger) Logf(level Level, template string, args ...interface{}) {
	switch level {
	case DebugLevel:
		l.logger.Debugf(template, args...)
	case InfoLevel:
		l.logger.Infof(template, args...)
	case WarnLevel:
		l.logger.Warnf(template, args...)
	case ErrorLevel:
		l.logger.Errorf(template, args...)
	default:
		l.logger.Infof(template, args...)
	}
}

func (l *logger) Logw(level Level, msg string, keysAndValues ...interface{}) {
	switch level {
	case DebugLevel:
		l.logger.Debugw(msg, keysAndValues...)
	case InfoLevel:
		l.logger.Infow(msg, keysAndValues...)
	case WarnLevel:
		l.logger.Warnw(msg, keysAndValues...)
	case ErrorLevel:
		l.logger.Errorw(msg, keysAndValues...)
	default:
		l.logger.Infow(msg, keysAndValues...)
	}
}

func (l *logger) Sync() error {
	return l.logger.Sync()
}
