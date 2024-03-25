package ui

import (
	"errors"

	"github.com/maladroitthief/entree/pkg/engine/core"
	"github.com/maladroitthief/mosaic"
)

var (
	ErrUnboundKey   = errors.New("key is unbound")
	ErrNilKeyboard  = errors.New("keyboard is nil")
	ErrInputRepoNil = errors.New("input repo is nil")

	DefaultKeyBindings = map[core.Input]string{
		core.InputAccept:        "Enter",
		core.InputCancel:        "Backspace",
		core.InputMoveUp:        "W",
		core.InputMoveDown:      "S",
		core.InputMoveLeft:      "A",
		core.InputMoveRight:     "D",
		core.InputAttack:        "MouseButtonLeft",
		core.InputAttackUp:      "ArrowUp",
		core.InputAttackDown:    "ArrowDown",
		core.InputAttackLeft:    "ArrowLeft",
		core.InputAttackRight:   "ArrowRight",
		core.InputDodge:         "Space",
		core.InputInteract:      "E",
		core.InputUseItem:       "MouseButtonRight",
		core.InputUseConsumable: "Q",
		core.InputMap:           "M",
		core.InputMenu:          "Escape",
	}
)

type InputHandler struct {
	repo     InputRepository
	settings InputSettings

	currentKeys   []string
	currentCursor mosaic.Vector
	currentInputs []core.Input
	inputStates   map[core.Input]int
}

type InputState struct {
	Cursor mosaic.Vector
	Keys   []string
}

type InputSettings struct {
	Keyboard map[core.Input]string
}

type InputRepository interface {
	GetInputSettings() (InputSettings, error)
	SetInputSettings(InputSettings) error
}

func NewInputHandler(r InputRepository) (*InputHandler, error) {
	if r == nil {
		return nil, ErrInputRepoNil
	}

	h := &InputHandler{
		repo: r,
	}

	err := h.Load()
	if err != nil {
		return nil, err
	}

	return h, nil
}

func (h *InputHandler) Update(state InputState) error {
	if h.inputStates == nil {
		h.inputStates = map[core.Input]int{}
	}

	h.currentKeys = state.Keys
	h.currentInputs = make([]core.Input, len(h.currentKeys))
	h.currentCursor = state.Cursor

	for i, k := range h.settings.Keyboard {
		for _, arg := range state.Keys {
			if k == arg {
				h.currentInputs = append(h.currentInputs, i)
				h.inputStates[i]++
				continue
			}

			h.inputStates[i] = 0
		}
	}

	return nil
}

func (h *InputHandler) Load() error {
	is, err := h.repo.GetInputSettings()
	if err != nil {
		return err
	}

	h.settings = is
	return nil
}

func (h *InputHandler) IsAny() bool {
	return len(h.currentKeys) > 0
}

func (h *InputHandler) IsPressed(i core.Input) bool {
	return h.inputStates[i] >= 1
}

func (h *InputHandler) IsJustPressed(i core.Input) bool {
	return h.inputStates[i] == 1
}

func (h *InputHandler) GetCursor() mosaic.Vector {
	return h.currentCursor
}

func (h *InputHandler) CurrentInputs() []core.Input {
	return h.currentInputs
}

func (i *InputSettings) Validate() error {
	if i.Keyboard == nil {
		return ErrNilKeyboard
	}

	return nil
}

func DefaultInputSettings() InputSettings {
	return InputSettings{
		Keyboard: DefaultKeyBindings,
	}
}
