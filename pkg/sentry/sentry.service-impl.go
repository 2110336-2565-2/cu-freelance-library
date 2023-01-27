package sentry

import (
	"github.com/getsentry/sentry-go"
	"time"
)

type service struct {
}

func NewService(DSN string) (Service, error) {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              DSN,
		TracesSampleRate: 1.0,
	}); err != nil {
		return nil, err
	}

	return &service{}, nil
}

func (s *service) CaptureMessage(msg string) {
	defer sentry.Flush(2 * time.Second)
	
	sentry.CaptureMessage(msg)
}

func (s *service) CaptureException(err error) {
	defer sentry.Flush(2 * time.Second)

	sentry.CaptureException(err)
}
