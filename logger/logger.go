package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log logger
)

// Logger specifies the interface for all log operations.
type Logger interface {
	Printf(string, ...interface{})
}

type logger struct {
	log *zap.Logger
}

func init() {
	conf := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
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

func (l logger) Printf(format string, v ...interface{}) {
	if len(v) == 0 {
		l.log.Info(format)
	} else {
		l.log.Info(fmt.Sprintf(format, v...))
	}
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
