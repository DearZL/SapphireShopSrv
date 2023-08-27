package config

import (
	"go.uber.org/zap"
)

func ZapConfig(way bool) *zap.Logger {
	var logger *zap.Logger
	var err error
	if way {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		panic(err)
	}
	return logger
}

func SyncLog(logger *zap.Logger) {
	err := logger.Sync()
	if err != nil {
		logger.Info("Log刷新缓冲区出错")
		return
	}
}
