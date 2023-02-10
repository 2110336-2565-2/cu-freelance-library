package sentry

type Sentry interface {
	CaptureMessage(msg string)
	CaptureException(err error)
}
