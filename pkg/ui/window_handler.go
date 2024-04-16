package ui

import (
	"errors"
)

var (
	ErrBlankTitle       = errors.New("title cannot be blank")
	ErrWindowWidthZero  = errors.New("window width cannot be zero")
	ErrWindowHeightZero = errors.New("window height cannot be zero")
	ErrWindowRepoNil    = errors.New("window repo is nil")
)

type WindowHandler struct {
	repo     WindowRepository
	settings WindowSettings
}

type WindowSettings struct {
	Width  int
	Height int
	Title  string
}

type WindowRepository interface {
	GetWindowSettings() (WindowSettings, error)
	SetWindowSettings(WindowSettings) error
}

func DefaultWindowSettings() WindowSettings {
	return WindowSettings{
		Width:  1920,
		Height: 1080,
		Title:  "Entree",
	}
}

func NewWindowHandler(r WindowRepository) (*WindowHandler, error) {
	if r == nil {
		return nil, ErrWindowRepoNil
	}

	h := &WindowHandler{
		repo: r,
	}

	err := h.Load()
	if err != nil {
		return nil, err
	}

	return h, nil
}

func (svc *WindowHandler) Height() int {
	return svc.settings.Height
}

func (svc *WindowHandler) Width() int {
	return svc.settings.Width
}

func (svc *WindowHandler) Title() string {
	return svc.settings.Title
}

func (svc *WindowHandler) Load() error {
	ws, err := svc.repo.GetWindowSettings()
	if err != nil {
		return err
	}

	svc.settings = ws

	return nil
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

	return nil
}
