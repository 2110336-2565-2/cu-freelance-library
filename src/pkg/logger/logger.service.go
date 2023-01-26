package logger

type Service interface {
	SetName(name string)
	Info(msg string, options ...interface{})
	Error(msg string, err error, options ...interface{})
	Debug(msg string, options ...interface{})
	Fatal(msg string, options ...interface{})
}
