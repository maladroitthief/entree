package settings

import "errors"

var (
	ErrBlankTitle       = errors.New("title cannot be blank")
	ErrWindowWidthZero  = errors.New("window width cannot be zero")
	ErrWindowHeightZero = errors.New("window height cannot be zero")
	ErrScaleZero        = errors.New("window scale cannot be zero")
)

type WindowSettings struct {
	Width  int
	Height int
	Title  string
	Scale  float64
}

func DefaultWindowSettings() WindowSettings {
	return WindowSettings{
		Width:  1920,
		Height: 1080,
		Title:  "Entree",
		Scale:  1.0,
	}
}

func (w *WindowSettings) Validate() error {
	if w.Title == "" {
		return ErrBlankTitle
	}

	if w.Width <= 0 {
		return ErrWindowWidthZero
	}

	if w.Height <= 0 {
		return ErrWindowHeightZero
	}

	if w.Scale <= 0 {
		return ErrScaleZero
	}

	return nil
}
