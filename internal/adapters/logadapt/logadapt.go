package logadapt

import "github.com/yawnak/fuel-record-crud/pkg/logger"

func NewLoggerAdapter() logger.Logger {
	return logger.NewLogger()
}
