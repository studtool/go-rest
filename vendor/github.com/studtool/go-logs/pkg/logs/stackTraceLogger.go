package logs

import (
	"runtime/debug"

	"github.com/sirupsen/logrus"

	"github.com/studtool/go-conv/pkg/conv"
)

type StackTraceLogger struct {
	backend *logrus.Logger
	params  StackTraceLoggerParams
}

type StackTraceLoggerParams struct {
	CommonLoggerParams
}

func NewStackTraceLogger(params StackTraceLoggerParams) *StackTraceLogger {
	return &StackTraceLogger{
		params: params,
		backend: func() *logrus.Logger {
			log := logrus.StandardLogger()
			log.SetFormatter(&logrus.JSONFormatter{})
			return log
		}(),
	}
}

func (logger *StackTraceLogger) Debug(args ...interface{}) {
	info := logger.makeInfo()
	info["stack"] = conv.BytesToString(debug.Stack())

	logger.backend.WithFields(info).Debug(args...)
}

func (logger *StackTraceLogger) Debugf(format string, args ...interface{}) {
	info := logger.makeInfo()
	info["stack"] = conv.BytesToString(debug.Stack())

	logger.backend.WithFields(info).Debugf(format, args...)
}

func (logger *StackTraceLogger) Info(args ...interface{}) {
	info := logger.makeInfo()
	info["stack"] = conv.BytesToString(debug.Stack())

	logger.backend.WithFields(info).Info(args...)
}

func (logger *StackTraceLogger) Infof(format string, args ...interface{}) {
	info := logger.makeInfo()
	info["stack"] = conv.BytesToString(debug.Stack())

	logger.backend.WithFields(info).Infof(format, args...)
}

func (logger *StackTraceLogger) Warning(args ...interface{}) {
	info := logger.makeInfo()
	info["stack"] = conv.BytesToString(debug.Stack())

	logger.backend.WithFields(info).Warn(args...)
}

func (logger *StackTraceLogger) Warningf(format string, args ...interface{}) {
	info := logger.makeInfo()
	info["stack"] = conv.BytesToString(debug.Stack())

	logger.backend.WithFields(info).Warningf(format, args...)
}

func (logger *StackTraceLogger) Error(args ...interface{}) {
	info := logger.makeInfo()
	info["stack"] = conv.BytesToString(debug.Stack())

	logger.backend.WithFields(info).Error(args...)
}

func (logger *StackTraceLogger) Errorf(format string, args ...interface{}) {
	info := logger.makeInfo()
	info["stack"] = conv.BytesToString(debug.Stack())

	logger.backend.WithFields(info).Errorf(format, args...)
}

func (logger *StackTraceLogger) Fatal(args ...interface{}) {
	info := logger.makeInfo()
	info["stack"] = conv.BytesToString(debug.Stack())

	logger.backend.WithFields(info).Fatal(args...)
}

func (logger *StackTraceLogger) Fatalf(format string, args ...interface{}) {
	info := logger.makeInfo()
	info["stack"] = conv.BytesToString(debug.Stack())

	logger.backend.WithFields(info).Fatalf(format, args...)
}

func (logger *StackTraceLogger) makeInfo() logrus.Fields {
	return logrus.Fields{
		"pid":               logger.params.PID,
		"host":              logger.params.Host,
		"component_name":    logger.params.ComponentName,
		"component_version": logger.params.ComponentVersion,
		"commit_hash":       logger.params.CommitHash,
	}
}
