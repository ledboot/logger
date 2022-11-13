package logger

import (
	"context"
	"fmt"
	"time"

	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

type gormLoggerImpl struct {
	ctx           context.Context
	log           Logger
	logLvl        gormlogger.LogLevel
	slowThreshold time.Duration
}

// Logger return singleton logger
func NewGormLogger(ctx context.Context, log Logger, slowThreshold time.Duration) gormlogger.Interface {
	return &gormLoggerImpl{
		ctx:           ctx,
		log:           log,
		logLvl:        gormlogger.Info,
		slowThreshold: slowThreshold,
	}
}

func (l *gormLoggerImpl) LogMode(lvl gormlogger.LogLevel) gormlogger.Interface {
	l.logLvl = lvl
	// l.log.SetLevel(toLoggerLevel(lvl))
	return l
}

func (l *gormLoggerImpl) Info(ctx context.Context, format string, args ...interface{}) {
	l.log.Infof(format, args...)
}

func (l *gormLoggerImpl) Warn(ctx context.Context, format string, args ...interface{}) {
	l.log.Warnf(format, args...)
}

func (l *gormLoggerImpl) Error(ctx context.Context, format string, args ...interface{}) {
	l.log.Errorf(format, args...)
}

var (
	traceStr     = "%s\n[%.3fms] [rows:%v] %s"
	traceWarnStr = "%s %s\n[%.3fms] [rows:%v] %s"
	traceErrStr  = "%s %s\n[%.3fms] [rows:%v] %s"
)

// Trace print sql message
func (l *gormLoggerImpl) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.logLvl > gormlogger.Silent {
		elapsed := time.Since(begin)
		switch {
		case err != nil && l.logLvl >= gormlogger.Error:
			sql, rows := fc()
			if rows == -1 {
				l.log.Errorf(traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				l.log.Errorf(traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		case elapsed > l.slowThreshold && l.slowThreshold != 0 && l.logLvl >= gormlogger.Warn:
			sql, rows := fc()
			slowLog := fmt.Sprintf("SLOW SQL >= %v", l.slowThreshold)
			if rows == -1 {
				l.log.Warnf(traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				l.log.Warnf(traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		case l.logLvl == gormlogger.Info:
			sql, rows := fc()
			if rows == -1 {
				l.log.Infof(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				l.log.Infof(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		}
	}
}
