package logging

import (
	"context"
	"log"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

type LogLevel struct {
	level logger.LogLevel
}

func NewLogger() logger.Interface {
	return &LogLevel{
		level: logger.Info,
	}
}

func (l *LogLevel) Error(ctx context.Context, msg string, data ...any) {
	if l.level >= logger.Error {
		Logger.Error(msg)
	}
}

func (l *LogLevel) Info(ctx context.Context, msg string, data ...any) {
	if l.level >= logger.Info {
		Logger.Info(msg)
	}
}

func (l *LogLevel) Warn(ctx context.Context, msg string, data ...any) {
	if l.level >= logger.Warn {
		Logger.Warn(msg)
	}
}

func (l *LogLevel) LogMode(logLevel logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.level = logLevel
	return &newLogger
}

func (l *LogLevel) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	fields := []zap.Field{
		zap.String("caller", utils.FileWithLineNum()),
		zap.Duration("elapsed_time", elapsed),
		zap.String("sql", sql),
		zap.Int64("rows_affected", rows),
	}

	Logger.Info("SQL Tracking", fields...)
}

var Logger *zap.Logger

func init() {
	var err error
	Logger, err = zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
}
