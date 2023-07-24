package ui

import (
	"errors"

	"github.com/maladroitthief/entree/common/logs"
)

var (
	ErrBlankTitle       = errors.New("title cannot be blank")
	ErrWindowWidthZero  = errors.New("window width cannot be zero")
	ErrWindowHeightZero = errors.New("window height cannot be zero")
	ErrScaleZero        = errors.New("window scale cannot be zero")
	ErrWindowRepoNil    = errors.New("window repo is nil")
)

type WindowHandler struct {
	repo     WindowRepository
	log      logs.Logger
	settings WindowSettings
}

type WindowSettings struct {
	Width  int
	Height int
	Title  string
	Scale  float64
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
		Scale:  1.0,
	}
}

func NewWindowHandler(l logs.Logger, r WindowRepository) (*WindowHandler, error) {
	if l == nil {
		return nil, ErrLoggerNil
	}

	if r == nil {
		return nil, ErrWindowRepoNil
	}

	h := &WindowHandler{
		repo: r,
		log:  l,
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

func (svc *WindowHandler) Scale() float64 {
	return svc.settings.Scale
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

	if w.Scale <= 0 {
		return ErrScaleZero
	}

	return nil
}
