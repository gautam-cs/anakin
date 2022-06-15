package config

import (
	"context"
	"time"

	"github.com/phuslu/log"
	"gorm.io/gorm/logger"
)

//GormLogger logger for gorm
type GormLogger struct {
	Log           log.Logger
	SlowThreshold time.Duration
	Silent        bool
}

// LogMode impl
func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	var newLogger = GormLogger{Log: l.Log}
	switch level {
	case logger.Silent:
		newLogger.Silent = true
	case logger.Error:
		newLogger.Log.SetLevel(log.ErrorLevel)
	case logger.Warn:
		newLogger.Log.SetLevel(log.WarnLevel)
	case logger.Info:
		newLogger.Log.SetLevel(log.InfoLevel)
	}

	return &newLogger
}

// Info impl
func (l *GormLogger) Info(ctx context.Context, format string, args ...interface{}) {
	l.Log.Info().Msgf(format, args...)
}

// Warn impl
func (l *GormLogger) Warn(ctx context.Context, format string, args ...interface{}) {
	l.Log.Warn().Msgf(format, args...)
}

// Error impl
func (l *GormLogger) Error(ctx context.Context, format string, args ...interface{}) {
	l.Log.Error().Msgf(format, args...)
}

// Trace impl
func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.Silent {
		return
	}
	elapsed := time.Since(begin)
	switch {
	case err != nil && l.Log.Level >= log.ErrorLevel:
		sql, rows := fc()
		if rows == -1 {
			l.Log.Error().Caller(1).Err(err).Dur("elapsed", elapsed).Str("sql", sql).Msg("")
		} else {
			l.Log.Error().Caller(1).Err(err).Dur("elapsed", elapsed).Str("sql", sql).Int64("rows", rows).Msg("")
		}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.Log.Level >= log.WarnLevel:
		sql, rows := fc()
		if rows == -1 {
			l.Log.Warn().Caller(1).Err(err).Dur("elapsed", elapsed).Str("sql", sql).Msgf("SLOW SQL >= %v", l.SlowThreshold)
		} else {
			l.Log.Warn().Caller(1).Err(err).Dur("elapsed", elapsed).Str("sql", sql).Int64("rows", rows).Msgf("SLOW SQL >= %v", l.SlowThreshold)
		}
	case l.Log.Level == log.InfoLevel:
		sql, rows := fc()
		if rows == -1 {
			l.Log.Info().Caller(1).Err(err).Dur("elapsed", elapsed).Str("sql", sql).Msg("")
		} else {
			l.Log.Info().Caller(1).Err(err).Dur("elapsed", elapsed).Str("sql", sql).Int64("rows", rows).Msg("")
		}
	case l.Log.Level == log.DebugLevel:
		sql, rows := fc()
		if rows == -1 {
			l.Log.Debug().Caller(1).Err(err).Dur("elapsed", elapsed).Str("sql", sql).Msg("")
		} else {
			l.Log.Debug().Caller(1).Err(err).Dur("elapsed", elapsed).Str("sql", sql).Int64("rows", rows).Msg("")
		}
	}
}
