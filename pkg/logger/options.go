package logger

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

type OptFunc func(*Config)

func applyOption(config *Config) func(OptFunc, int) {
	return func(option OptFunc, index int) {
		option(config)
	}
}

type Config struct {
	Level   zerolog.Level
	Writers []io.Writer
}

var DefaultConfig = Config{
	Level: zerolog.InfoLevel,
}

func WithLevel(level zerolog.Level) OptFunc {
	return func(config *Config) {
		config.Level = level
	}
}

func WithWriter(writer io.Writer) OptFunc {
	return func(config *Config) {
		config.Writers = append(config.Writers, writer)
	}
}

func WithConsoleWriter() OptFunc {
	return func(config *Config) {
		config.Writers = append(config.Writers, zerolog.ConsoleWriter{Out: os.Stderr})
	}
}
