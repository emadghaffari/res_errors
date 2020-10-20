package logger

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	envLogLevel  = "LOG_LEVEL"
	envLogOutPut = "LOG_OUTPUT"
)

var (
	log logger
)

// Logger specifies the interface for all log operations.
type Logger interface {
	Print(...interface{})
	Printf(string, ...interface{})
}

type logger struct {
	log *zap.Logger
}

// GetLogger func
func GetLogger() Logger {
	return log
}

func init() {
	conf := zap.Config{
		OutputPaths: []string{getOutPut()},
		Level:       zap.NewAtomicLevelAt(getLevel()),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseColorLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	var err error
	if log.log, err = conf.Build(); err != nil {
		panic(err)
	}
}

func getLevel() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(envLogLevel))) {
	case "info":
		return zap.InfoLevel
	case "error":
		return zap.ErrorLevel
	case "debug":
		return zap.DebugLevel
	default:
		return zap.InfoLevel
	}
}

func getOutPut() string {
	output := strings.ToLower(strings.TrimSpace(os.Getenv(envLogOutPut)))
	if output == "" {
		return "stdout"

	}
	return output
}

func (l logger) Printf(format string, v ...interface{}) {
	if len(v) == 0 {
		l.log.Info(format)
	} else {
		l.log.Info(fmt.Sprintf(format, v...))
	}
}

func (l logger) Print(v ...interface{}) {
	l.log.Info(fmt.Sprintf("%v", v))
}

// Info func
func Info(msg string, tags ...zap.Field) {
	log.log.Info(msg, tags...)
	log.log.Sync()
}

// Error func
func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.log.Error(msg, tags...)
	log.log.Sync()
}
