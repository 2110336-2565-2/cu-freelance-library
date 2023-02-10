package loggersdk

import (
	"github.com/2110336-2565-2/sec3-group15-cu-freelance-logger.git/pkg/sentry"
	"go.uber.org/zap"
)

type service struct {
	serviceName string
	*zap.SugaredLogger
	sentryService sentry.Service
}

func NewService(serviceName string, dsn string) (Service, error) {
	lg, _ := zap.NewProduction()
	sugar := lg.Sugar()

	sentrySrv, err := sentry.NewService(dsn)
	if err != nil {
		return nil, err
	}

	return &service{serviceName, sugar, sentrySrv}, nil
}
func (s *service) SetName(name string) {
	s.serviceName = name
}

func (s *service) Info(msg string, options ...interface{}) {
	if len(options) > 0 {
		s.SugaredLogger.Infow(msg, "service_name", s.serviceName, options)
		return
	}

	s.SugaredLogger.Infow(msg, "service_name", s.serviceName)
}

func (s *service) Error(msg string, err error, options ...interface{}) {
	s.sentryService.CaptureException(err)
	if len(options) > 0 {
		s.SugaredLogger.Errorw(msg, "error", zap.Error(err), "service_name", s.serviceName, options)
		return
	}

	s.SugaredLogger.Errorw(msg, "error", zap.Error(err), "service_name", s.serviceName)
}

func (s *service) Debug(msg string, options ...interface{}) {
	if len(options) > 0 {
		s.SugaredLogger.Debugw(msg, "service_name", s.serviceName, options)
		return
	}

	s.SugaredLogger.Debugw(msg, "service_name", s.serviceName)
}

func (s *service) Fatal(msg string, options ...interface{}) {
	if len(options) > 0 {
		s.SugaredLogger.Fatalw(msg, "service_name", s.serviceName, options)
		return
	}

	s.SugaredLogger.Fatalw(msg, "service_name", s.serviceName)
}
