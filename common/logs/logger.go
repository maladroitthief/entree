package logs

type Logger interface {
  Info(message string)
  Error(methodName string, args interface{}, err error)
  Fatal(methodName string, args interface{}, err error)
}
