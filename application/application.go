package application

import "errors"

var (
	ErrLoggerNil          = errors.New("nil logger")
	ErrSettingsServiceNil = errors.New("nil settings service")
)

type Inputs struct {
	CursorX int
	CursorY int
	Inputs  []string
}
