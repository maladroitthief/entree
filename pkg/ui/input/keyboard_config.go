package input

import (
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type keyboardConfig struct {
	current      []action
	keys         map[action]ebiten.Key
	mouseButtons map[action]ebiten.MouseButton
}

var (
	ErrUnknownKey         = errors.New("the key does not exist")
	ErrUnknownMouseButton = errors.New("the mouse button does not exist")
)

func NewKeyboardConfig() *keyboardConfig {
  kbc := &keyboardConfig{}
  kbc.Initialize()

  return kbc
}

func DefaultKey(a action) (ebiten.Key, error) {
	switch a {
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

	return ebiten.KeyA, ErrUnknownKey
}

func DefaultMouseButton(a action) (ebiten.MouseButton, error) {
	switch a {
	case Ok:
		return ebiten.MouseButtonLeft, nil
	}

	return ebiten.MouseButtonLeft, ErrUnknownMouseButton
}

func (c *keyboardConfig) Initialize() {
	if c.keys == nil {
		c.keys = map[action]ebiten.Key{}
		c.ResetAllKeys()
	}

	if c.mouseButtons == nil {
		c.mouseButtons = map[action]ebiten.MouseButton{}
		c.ResetAllMouseButtons()
	}
}

func (c *keyboardConfig) ResetAllKeys() {
	for _, a := range actions {
		k, err := DefaultKey(a)

		if err != nil {
			continue
		}

		c.SetKey(a, k)
	}
}

func (c *keyboardConfig) SetKey(a action, k ebiten.Key) {
	c.keys[a] = k

	_, ok := c.mouseButtons[a]
	if ok {
		delete(c.mouseButtons, a)
	}
}

func (c *keyboardConfig) ResetAllMouseButtons() {
	for _, a := range actions {
		mb, err := DefaultMouseButton(a)

		if err != nil {
			continue
		}

		c.SetMouseButton(a, mb)
	}
}

func (c *keyboardConfig) SetMouseButton(a action, mb ebiten.MouseButton) {
	c.mouseButtons[a] = mb

	_, ok := c.keys[a]
	if ok {
		delete(c.keys, a)
	}
}

func (c *keyboardConfig) IsPressed(b action) bool {
	c.Initialize()

	k, ok := c.keys[b]
	if ok {
		return ebiten.IsKeyPressed(k)
	}

	mb, ok := c.mouseButtons[b]
	if ok {
		return ebiten.IsMouseButtonPressed(mb)
	}

	dk, err := DefaultKey(b)
	if err == nil {
		c.keys[b] = dk
		return ebiten.IsKeyPressed(dk)
	}

	dmb, err := DefaultMouseButton(b)
	if err == nil {
		c.mouseButtons[b] = dmb
		return ebiten.IsMouseButtonPressed(dmb)
	}

	return false
}

func (c *keyboardConfig) IsJustPressed(b action) bool {
	c.Initialize()

	k, ok := c.keys[b]
	if ok {
		return inpututil.IsKeyJustPressed(k)
	}

	mb, ok := c.mouseButtons[b]
	if ok {
		return inpututil.IsMouseButtonJustPressed(mb)
	}

	dk, err := DefaultKey(b)
	if err == nil {
		c.keys[b] = dk
		return inpututil.IsKeyJustPressed(dk)
	}

	dmb, err := DefaultMouseButton(b)
	if err == nil {
		c.mouseButtons[b] = dmb
		return inpututil.IsMouseButtonJustPressed(dmb)
	}

	return false
}
