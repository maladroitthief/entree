package logs

type Logger interface {
	SetLevel(level string)
	Debug(message string, context interface{})
	Info(message string, context interface{})
	Error(method string, context interface{}, err error)
	Fatal(method string, context interface{}, err error)
}
