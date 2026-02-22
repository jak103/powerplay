package db

import (
	"context"
	"errors"
	"time"

	"github.com/jak103/powerplay/internal/utils/log"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

// Interface logger interface
type dbLogger struct {
	theLogger *log.Logger
	level     logger.LogLevel
}

// LogMode log mode
func (l *dbLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.level = level
	return l
}

// Info print info
func (l dbLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.level >= logger.Info {
		l.theLogger.Info(msg, data...)
	}
}

// Warn print warn messages
func (l dbLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.level >= logger.Warn {
		l.theLogger.Warn(msg, data...)
	}
}

// Error print error messages
func (l dbLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.level >= logger.Error {
		l.theLogger.Error(msg, data...)
	}
}

// Trace print sql message
func (l dbLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.level <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.level >= logger.Error && !errors.Is(err, logger.ErrRecordNotFound):
		sql, rows := fc()
		if rows == -1 {
			l.theLogger.WithErr(err).Error("%s SQL: %s -- [%v]", utils.FileWithLineNum(), sql, elapsed)
		} else {
			l.theLogger.WithErr(err).Error("%s SQL: %s -- [%v], rows affected: %v", utils.FileWithLineNum(), sql, elapsed, rows)
		}
	case elapsed > 1*time.Second && l.level >= logger.Warn:
		sql, rows := fc()
		if rows == -1 {
			l.theLogger.Warn("%s SLOW SQL: %s -- [%v]", utils.FileWithLineNum(), sql, elapsed)
		} else {
			l.theLogger.Warn("%s SLOW SQL: %s -- [%v], rows affected: %v", utils.FileWithLineNum(), sql, elapsed, rows)
		}
	case l.level == logger.Info:
		sql, rows := fc()

		if rows == -1 {
			l.theLogger.Info("%s SQL: %s -- [%v]", utils.FileWithLineNum(), sql, elapsed)
		} else {
			l.theLogger.Info("%s SQL: %s -- [%v], rows affected: %v", utils.FileWithLineNum(), sql, elapsed, rows)
		}
	}
}
