package decorator

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

type CommandHandler[C any] interface {
	Handle(cmd C) error
}

func ApplyCommandDecorators[H any](
	handler CommandHandler[H],
	logger *logrus.Entry,
) CommandHandler[H] {
	return commandLoggingDecorator[H]{
    base: handler,
    logger: logger,
  }
}

func generateActionName(handler any) string {
	return strings.Split(fmt.Sprintf("%T", handler), ".")[1]
}
