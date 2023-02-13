package input

import (
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type keyboardConfig struct {
	current      []Input
	keys         map[Input]ebiten.Key
	mouseButtons map[Input]ebiten.MouseButton
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

func DefaultKey(i Input) (ebiten.Key, error) {
	switch i {
	case Accept:
		return ebiten.KeyEnter, nil
	case Cancel:
		return ebiten.KeyEscape, nil
	case MoveUp:
		return ebiten.KeyW, nil
	case MoveDown:
		return ebiten.KeyS, nil
	case MoveLeft:
		return ebiten.KeyA, nil
	case MoveRight:
		return ebiten.KeyD, nil
	}

	return ebiten.KeyA, ErrUnknownKey
}

func DefaultMouseButton(a Input) (ebiten.MouseButton, error) {
	switch a {
	case Accept:
		return ebiten.MouseButtonLeft, nil
	}

	return ebiten.MouseButtonLeft, ErrUnknownMouseButton
}

func (c *keyboardConfig) Initialize() {
	if c.keys == nil {
		c.keys = map[Input]ebiten.Key{}
		c.ResetAllKeys()
	}

	if c.mouseButtons == nil {
		c.mouseButtons = map[Input]ebiten.MouseButton{}
		c.ResetAllMouseButtons()
	}
}

func (c *keyboardConfig) ResetAllKeys() {
	for _, a := range Inputs() {
		k, err := DefaultKey(a)

		if err != nil {
			continue
		}

		c.SetKey(a, k)
	}
}

func (c *keyboardConfig) SetKey(i Input, k ebiten.Key) {
	c.keys[i] = k

	_, ok := c.mouseButtons[i]
	if ok {
		delete(c.mouseButtons, i)
	}
}

func (c *keyboardConfig) ResetAllMouseButtons() {
	for _, i := range Inputs() {
		mb, err := DefaultMouseButton(i)

		if err != nil {
			continue
		}

		c.SetMouseButton(i, mb)
	}
}

func (c *keyboardConfig) SetMouseButton(i Input, mb ebiten.MouseButton) {
	c.mouseButtons[i] = mb

	_, ok := c.keys[i]
	if ok {
		delete(c.keys, i)
	}
}

func (c *keyboardConfig) IsPressed(i Input) bool {
	c.Initialize()

	k, ok := c.keys[i]
	if ok {
		return ebiten.IsKeyPressed(k)
	}

	mb, ok := c.mouseButtons[i]
	if ok {
		return ebiten.IsMouseButtonPressed(mb)
	}

	dk, err := DefaultKey(i)
	if err == nil {
		c.keys[i] = dk
		return ebiten.IsKeyPressed(dk)
	}

	dmb, err := DefaultMouseButton(i)
	if err == nil {
		c.mouseButtons[i] = dmb
		return ebiten.IsMouseButtonPressed(dmb)
	}

	return false
}

func (c *keyboardConfig) IsJustPressed(i Input) bool {
	c.Initialize()

	k, ok := c.keys[i]
	if ok {
		return inpututil.IsKeyJustPressed(k)
	}

	mb, ok := c.mouseButtons[i]
	if ok {
		return inpututil.IsMouseButtonJustPressed(mb)
	}

	dk, err := DefaultKey(i)
	if err == nil {
		c.keys[i] = dk
		return inpututil.IsKeyJustPressed(dk)
	}

	dmb, err := DefaultMouseButton(i)
	if err == nil {
		c.mouseButtons[i] = dmb
		return inpututil.IsMouseButtonJustPressed(dmb)
	}

	return false
}

func (c *keyboardConfig) IsAnyKey() bool {
	pressedKeys := inpututil.AppendPressedKeys([]ebiten.Key{})

	if len(pressedKeys) > 0 {
		return true
	}

	return false
}
