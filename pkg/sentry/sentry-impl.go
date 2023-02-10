package sentry

import (
	sentrygo "github.com/getsentry/sentry-go"
	"time"
)

type sentry struct {
}

func NewService(DSN string) (Sentry, error) {
	if err := sentrygo.Init(sentrygo.ClientOptions{
		Dsn:              DSN,
		TracesSampleRate: 1.0,
	}); err != nil {
		return nil, err
	}

	return &sentry{}, nil
}

func (s *sentry) CaptureMessage(msg string) {
	defer sentrygo.Flush(2 * time.Second)

	sentrygo.CaptureMessage(msg)
}

func (s *sentry) CaptureException(err error) {
	defer sentrygo.Flush(2 * time.Second)

	sentrygo.CaptureException(err)
}
