package logger

import (
	"github.com/rs/zerolog"
	"github.com/samber/lo"
)

type Logger struct {
	logger zerolog.Logger
}

func NewLogger(options ...OptFunc) Logger {
	config := DefaultConfig
	lo.ForEach(options, applyOption(&config))
	multi := zerolog.MultiLevelWriter()

	newLogger := zerolog.New(multi)
	return Logger{newLogger}
}
