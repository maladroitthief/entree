package ui

import (
	"errors"
	"image"
	"image/color"

	"github.com/maladroitthief/entree/common/logs"
	"github.com/maladroitthief/entree/common/theme"
	"github.com/maladroitthief/entree/pkg/content"
	"github.com/maladroitthief/entree/pkg/engine/core"
)

const (
	TransitionMaxCount = 20
)

var (
	Termination = errors.New("game closed normally")

	ErrGraphicsServerNil = errors.New("graphics server is nil")
	ErrInputHandlerNil   = errors.New("input handler is nil")
	ErrWindowHandlerNil  = errors.New("window handler is nil")
)

type Scene interface {
	Update(*SceneState) error
	Size() (width, height int)
	CellSize() int
	GetState() *core.ECS
	GetCamera() *Camera
	BackgroundColor() color.Color
}

type SceneState struct {
	log   logs.Logger
	mgr   *SceneManager
	input *InputHandler
	theme theme.Colors
}

type SceneManager struct {
	log             logs.Logger
	currentScene    Scene
	nextScene       Scene
	transitionCount int
	theme           theme.Colors

	input    *InputHandler
	graphics *GraphicsServer
	window   *WindowHandler
}

func NewSceneManager(
	l logs.Logger,
	g *GraphicsServer,
	i *InputHandler,
	w *WindowHandler,
) (*SceneManager, error) {
	if l == nil {
		return nil, ErrLoggerNil
	}

	if g == nil {
		return nil, ErrGraphicsServerNil
	}

	if i == nil {
		return nil, ErrInputHandlerNil
	}

	if w == nil {
		return nil, ErrWindowHandlerNil
	}

	m := &SceneManager{
		log:      l,
		graphics: g,
		input:    i,
		window:   w,
		theme:    &theme.Endesga32{},
	}

	err := m.GoTo(NewTitleScene(m.sceneState()))

	return m, err
}

func (m *SceneManager) Update(state InputState) error {
	// Update Settings
	err := m.input.Update(state)
	if err != nil {
		return err
	}

	if m.currentScene == nil {
		err = m.GoTo(NewTitleScene(m.sceneState()))
	}

	if err != nil {
		return err
	}

	if m.transitionCount <= 0 {
		return m.currentScene.Update(m.sceneState())
	}

	m.transitionCount--

	if m.transitionCount > 0 {
		return nil
	}

	m.currentScene = m.nextScene
	m.nextScene = nil

	return nil
}

func (m *SceneManager) sceneState() *SceneState {
	return &SceneState{
		log:   m.log,
		mgr:   m,
		input: m.input,
		theme: m.theme,
	}
}

func (m *SceneManager) GetCamera() *Camera {
	return m.currentScene.GetCamera()
}

func (m *SceneManager) Size() (width, height int) {
	return m.currentScene.Size()
}

func (m *SceneManager) CellSize() int {
	return m.currentScene.CellSize()
}

func (m *SceneManager) GetState() *core.ECS {
	return m.currentScene.GetState()
}

func (m *SceneManager) BackgroundColor() color.Color {
	return m.currentScene.BackgroundColor()
}

func (m *SceneManager) GoTo(s Scene) error {
	if m.currentScene == nil {
		m.currentScene = s
	} else {
		m.nextScene = s
		m.transitionCount = TransitionMaxCount
	}

	return nil
}

func (m *SceneManager) Layout(width, height int) (screenWidth, screenHeight int) {
	return m.WindowSize()
}

func (m *SceneManager) WindowSize() (screenWidth, screenHeight int) {
	return m.window.Width(), m.window.Height()
}

func (m *SceneManager) WindowTitle() string {
	return m.window.Title()
}

func (m *SceneManager) Scale() float64 {
	return m.window.Scale()
}

func (m *SceneManager) SpriteSheet(sheet string) (*content.SpriteSheet, error) {
	return m.graphics.SpriteSheet(sheet)
}

func (m *SceneManager) SpriteRectangle(
	sheet string,
	sprite string,
) (image.Rectangle, error) {
	return m.graphics.Sprite(sheet, sprite)
}
