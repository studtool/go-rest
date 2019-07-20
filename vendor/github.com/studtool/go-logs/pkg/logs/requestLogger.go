package logs

import (
	"time"

	"github.com/sirupsen/logrus"
)

type RequestLogger struct {
	logger *logrus.Logger
	params RequestLoggerParams
}

type RequestLoggerParams struct {
	CommonLoggerParams
}

func NewRequestLogger(params RequestLoggerParams) *RequestLogger {
	return &RequestLogger{
		params: params,
		logger: func() *logrus.Logger {
			log := logrus.StandardLogger()
			log.SetFormatter(&logrus.JSONFormatter{})
			return log
		}(),
	}
}

type RequestParams struct {
	Method      string
	Path        string
	Status      int
	Type        string
	UserID      string
	ClientID    string
	RequestIP   string
	XRealIP     string
	ACLGroup    string
	UserAgent   string
	RequestTime time.Duration
}

const (
	RequestHandledMessage    = "request handled"
	RequestNotHandledMessage = "request not handled"
)

func (logger *RequestLogger) Info(p *RequestParams) {
	logger.logger.WithFields(logger.makeLogFields(p)).Info(RequestHandledMessage)
}

func (logger *RequestLogger) Warning(p *RequestParams) {
	logger.logger.WithFields(logger.makeLogFields(p)).Warning(RequestHandledMessage)
}

func (logger *RequestLogger) Error(p *RequestParams) {
	logger.logger.WithFields(logger.makeLogFields(p)).Error(RequestNotHandledMessage)
}

func (logger *RequestLogger) makeLogFields(p *RequestParams) logrus.Fields {
	return logrus.Fields{
		"pid":               logger.params.PID,
		"host":              logger.params.Host,
		"component_name":    logger.params.ComponentName,
		"component_version": logger.params.ComponentVersion,
		"commit_hash":       logger.params.CommitHash,
		"method":            p.Method,
		"path":              p.Path,
		"status":            p.Status,
		"type":              p.Type,
		"user_id":           p.UserID,
		"client_id":         p.ClientID,
		"request_ip":        p.RequestIP,
		"x_real_ip":         p.XRealIP,
		"acl_group":         p.ACLGroup,
		"user_agent":        p.UserAgent,
		"request_time":      p.RequestTime,
	}
}
