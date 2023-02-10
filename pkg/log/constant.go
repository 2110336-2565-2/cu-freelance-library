package log

type Level uint8

const (
	Fatal   Level = 13
	Error         = 8
	Warn          = 5
	Info          = 3
	Debug         = 1
	Unknown       = 0
)
