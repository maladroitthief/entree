package input

import (
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type keyboardConfig struct {
	current      []virtualGamepadButton
	buttons      map[virtualGamepadButton]ebiten.Key
	mouseButtons map[virtualGamepadButton]ebiten.MouseButton
}

var (
	ErrUnknownButton = errors.New("the button does not exist")
	ErrUnknownMouseButton = errors.New("the mouse button does not exist")
)

func DefaultKey(b virtualGamepadButton) (ebiten.Key, error) {
	switch b {
	case Ok:
		return ebiten.KeySpace, nil
	case Quit:
		return ebiten.KeyEscape, nil
	case Up:
		return ebiten.KeyW, nil
	case Down:
		return ebiten.KeyS, nil
	case Left:
		return ebiten.KeyA, nil
	case Right:
		return ebiten.KeyD, nil
	}

	return ebiten.KeyA, ErrUnknownButton
}

func DefaultMouseButton(b virtualGamepadButton) (ebiten.MouseButton, error) {
	switch b {
	case Ok:
		return ebiten.MouseButtonLeft, nil
	}

	return ebiten.MouseButtonLeft, ErrUnknownMouseButton
}

func (c *keyboardConfig) initialize() {
	if c.buttons == nil {
		c.buttons = map[virtualGamepadButton]ebiten.Key{}
	}

	if c.mouseButtons == nil {
		c.mouseButtons = map[virtualGamepadButton]ebiten.MouseButton{}
	}
}

func (c *keyboardConfig) Reset() {
	c.buttons = nil
}

func (c *keyboardConfig) IsButtonPressed(b virtualGamepadButton) bool {
	c.initialize()

	k, ok := c.buttons[b]
	if ok {
		return ebiten.IsKeyPressed(k)
	}

	mb, ok := c.mouseButtons[b]
	if ok {
		return ebiten.IsMouseButtonPressed(mb)
	}

  dk, err := DefaultKey(b)
  if err == nil {
    c.buttons[b] = dk
    return ebiten.IsKeyPressed(dk)
  }

  dmb, err := DefaultMouseButton(b)
  if err == nil {
    c.mouseButtons[b] = dmb
		return ebiten.IsMouseButtonPressed(dmb)
  }

	return false
}

func (c *keyboardConfig) IsButtonJustPressed(b virtualGamepadButton) bool {
	c.initialize()

	k, ok := c.buttons[b]
	if ok {
		return inpututil.IsKeyJustPressed(k)
	}

	mb, ok := c.mouseButtons[b]
	if ok {
		return inpututil.IsMouseButtonJustPressed(mb)
	}

  dk, err := DefaultKey(b)
  if err == nil {
    c.buttons[b] = dk
		return inpututil.IsKeyJustPressed(dk)
  }

  dmb, err := DefaultMouseButton(b)
  if err == nil {
    c.mouseButtons[b] = dmb
		return inpututil.IsMouseButtonJustPressed(dmb)
  }

	return false
}
