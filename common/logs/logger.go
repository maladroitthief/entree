package logs

type Logger interface {
  Info(message string, args interface{})
  Error(methodName string, args interface{}, err error)
  Fatal(methodName string, args interface{}, err error)
}
