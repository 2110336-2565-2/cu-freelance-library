package sentry

type Service interface {
	CaptureMessage(msg string)
	CaptureException(err error)
}
