package gosdk

import (
	"github.com/2110336-2565-2/cu-freelance-library/pkg/log"
	"github.com/2110336-2565-2/cu-freelance-library/pkg/sentry"
	"go.uber.org/zap"
	"time"
)

type Logger struct {
	keyword     map[string]any
	level       log.Level
	serviceName string
	err         error
	*zap.Logger
	sentryService sentry.Sentry
}

func NewLogger(serviceName string) *Logger {
	lg, _ := zap.NewProduction()

	keyword := map[string]any{}

	return &Logger{
		keyword,
		log.Unknown,
		serviceName,
		nil,
		lg,
		nil,
	}
}
func (s *Logger) SetSentryDSN(dsn string) error {
	sentrySrv, err := sentry.NewService(dsn)
	if err != nil {
		return err
	}

	s.sentryService = sentrySrv
	return nil
}

func (s *Logger) SetName(name string) {
	s.serviceName = name
}

func (s *Logger) Info() *Logger {
	s.level = log.Info
	return s
}

func (s *Logger) Error(err error) *Logger {
	s.level = log.Error
	s.err = err
	return s
}

func (s *Logger) Warn() *Logger {
	s.level = log.Warn
	return s
}

func (s *Logger) Debug() *Logger {
	s.level = log.Debug
	return s
}

func (s *Logger) Fatal(err error) *Logger {
	s.level = log.Fatal
	s.err = err
	return s
}

func (s *Logger) Keyword(key string, val any) *Logger {
	s.keyword[key] = val
	return s
}

func (s *Logger) Msg(msg string) {
	defer s.Logger.Sync()

	var keywordList []zap.Field
	s.createKeywordList(&keywordList)

	switch s.level {
	case log.Debug:
		s.Logger.Debug(msg, keywordList...)
	case log.Info:
		s.Logger.Info(msg, keywordList...)
	case log.Warn:
		s.Logger.Warn(msg, keywordList...)
	case log.Error:
		if s.sentryService != nil {
			s.sentryService.CaptureException(s.err)
		}
		s.Logger.Error(msg, keywordList...)
	case log.Fatal:
		if s.sentryService != nil {
			s.sentryService.CaptureException(s.err)
		}
		s.Logger.Fatal(msg, keywordList...)
	default:
		return
	}

	s.keyword = map[string]any{}
	s.err = nil
}

func (s *Logger) createKeywordList(keywords *[]zap.Field) {
	*keywords = append(*keywords, zap.String("service", s.serviceName))

	if s.err != nil {
		*keywords = append(*keywords, zap.Error(s.err))
	}

	for k, v := range s.keyword {
		switch v.(type) {
		case string:
			*keywords = append(*keywords, zap.String(k, v.(string)))
		case int:
			*keywords = append(*keywords, zap.Int(k, v.(int)))
		case time.Duration:
			*keywords = append(*keywords, zap.Duration(k, v.(time.Duration)))
		}

	}
}
