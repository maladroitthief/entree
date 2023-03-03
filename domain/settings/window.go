package settings

import "errors"

var (
	ErrBlankTitle = errors.New("title cannot be blank")
	ErrWindowWidthZero = errors.New("window width cannot be zero")
	ErrWindowHeightZero = errors.New("window height cannot be zero")
)

type Window struct {
	Width  int
	Height int
	Title  string
}

func WindowDefaults() Window {
	return Window{
		Width:  1920,
		Height: 1080,
		Title:  "Entree",
	}
}

func (w *Window) Validate() error {
	if w.Title == "" {
    return ErrBlankTitle
	}

	if w.Width <= 0 {
    return ErrWindowWidthZero
	}

	if w.Height <= 0 {
    return ErrWindowHeightZero
	}

	return nil
}
