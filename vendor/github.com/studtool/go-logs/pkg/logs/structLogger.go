package logs

import (
	"github.com/sirupsen/logrus"
)

type StructLogger struct {
	entry *logrus.Entry
}

type StructLoggerParams struct {
	CommonLoggerParams
	StructWithPkgName string
}

func NewStructLogger(params StructLoggerParams) *StructLogger {
	backend := logrus.StandardLogger()
	backend.SetFormatter(&logrus.JSONFormatter{})

	fields := logrus.Fields{
		"pid":               params.PID,
		"host":              params.Host,
		"component_name":    params.ComponentName,
		"component_version": params.ComponentVersion,
		"commit_hash":       params.CommitHash,
		"struct":            params.StructWithPkgName,
	}

	return &StructLogger{
		entry: backend.WithFields(fields),
	}
}

func (logger *StructLogger) Debug(args ...interface{}) {
	logger.entry.Debug(args...)
}

func (logger *StructLogger) Debugf(format string, args ...interface{}) {
	logger.entry.Debugf(format, args...)
}

func (logger *StructLogger) Info(args ...interface{}) {
	logger.entry.Info(args...)
}

func (logger *StructLogger) Infof(format string, args ...interface{}) {
	logger.entry.Infof(format, args...)
}

func (logger *StructLogger) Warning(args ...interface{}) {
	logger.entry.Warn(args...)
}

func (logger *StructLogger) Warningf(format string, args ...interface{}) {
	logger.entry.Warningf(format, args...)
}

func (logger *StructLogger) Error(args ...interface{}) {
	logger.entry.Error(args...)
}

func (logger *StructLogger) Errorf(format string, args ...interface{}) {
	logger.entry.Errorf(format, args...)
}

func (logger *StructLogger) Fatal(args ...interface{}) {
	logger.entry.Fatal(args...)
}

func (logger *StructLogger) Fatalf(format string, args ...interface{}) {
	logger.entry.Fatalf(format, args...)
}
