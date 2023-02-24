package gosdk

import (
	"github.com/2110336-2565-2/cu-freelance-library/pkg/log"
	"github.com/2110336-2565-2/cu-freelance-library/pkg/sentry"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var sentryService sentry.Sentry

type Logger interface {
	SetName(name string)
	Fatal(err error) Logger
	Error(err error) Logger
	Warn() Logger
	Info() Logger
	Debug() Logger
	Keyword(key string, val any) Logger
	Msg(msg string)
}

type logger struct {
	keyword     map[string]any
	level       log.Level
	serviceName string
	err         error
	*zap.Logger
}

func NewLogger(serviceName string) Logger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339Nano)
	config.EncoderConfig.TimeKey = "time"

	lg, _ := config.Build()

	keyword := map[string]any{}

	return &logger{
		keyword,
		log.Unknown,
		serviceName,
		nil,
		lg,
	}
}
func SetSentryDSN(newDSN string) error {
	sentrySrv, err := sentry.NewService(newDSN)
	if err != nil {
		return err
	}

	sentryService = sentrySrv
	return nil
}

func (s *logger) SetName(name string) {
	s.serviceName = name
}

func (s *logger) Info() Logger {
	s.level = log.Info
	return s
}

func (s *logger) Error(err error) Logger {
	s.level = log.Error
	s.err = err
	return s
}

func (s *logger) Warn() Logger {
	s.level = log.Warn
	return s
}

func (s *logger) Debug() Logger {
	s.level = log.Debug
	return s
}

func (s *logger) Fatal(err error) Logger {
	s.level = log.Fatal
	s.err = err
	return s
}

func (s *logger) Keyword(key string, val any) Logger {
	s.keyword[key] = val
	return s
}

func (s *logger) Msg(msg string) {
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
		if sentryService != nil {
			sentryService.CaptureException(s.err)
		}
		s.Logger.Error(msg, keywordList...)
	case log.Fatal:
		if sentryService != nil {
			sentryService.CaptureException(s.err)
		}
		s.Logger.Fatal(msg, keywordList...)
	default:
		return
	}

	s.keyword = map[string]any{}
	s.err = nil
}

func (s *logger) createKeywordList(keywords *[]zap.Field) {
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
		case bool:
			*keywords = append(*keywords, zap.Bool(k, v.(bool)))
		default:
			*keywords = append(*keywords, zap.Any(k, v.(any)))
		}

	}
}
