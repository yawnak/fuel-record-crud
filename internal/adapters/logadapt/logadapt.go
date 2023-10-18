package logadapt

import (
	"github.com/rs/zerolog"
	"github.com/yawnak/fuel-record-crud/pkg/logger"
)

func NewStagingLoggerAdapter() logger.Logger {
	return logger.NewLogger()
}

func NewProductionLoggerAdapter() logger.Logger {
	return logger.NewLogger()
}

func NewDevelopmentLoggerAdapter() logger.Logger {
	return logger.NewLogger(logger.WithConsoleWriter(), logger.WithLevel(zerolog.DebugLevel))
}
