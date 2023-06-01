package application_test

type logger struct {
}

func (l *logger) Info(message string, args interface{}) {}

func (l *logger) Error(methodName string, args interface{}, err error) {}

func (l *logger) Fatal(methodName string, args interface{}, err error) {}
