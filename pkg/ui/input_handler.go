package ui

import (
	"errors"

	"github.com/maladroitthief/entree/common/data"
	"github.com/maladroitthief/entree/common/logs"
	"github.com/maladroitthief/entree/pkg/engine/core"
)

var (
	ErrUnboundKey   = errors.New("key is unbound")
	ErrNilKeyboard  = errors.New("keyboard is nil")
	ErrInputRepoNil = errors.New("input repo is nil")

	DefaultKeyBindings = map[core.Input]string{
		core.Accept:        "Enter",
		core.Cancel:        "Backspace",
		core.MoveUp:        "W",
		core.MoveDown:      "S",
		core.MoveLeft:      "A",
		core.MoveRight:     "D",
		core.Attack:        "MouseButtonLeft",
		core.AttackUp:      "ArrowUp",
		core.AttackDown:    "ArrowDown",
		core.AttackLeft:    "ArrowLeft",
		core.AttackRight:   "ArrowRight",
		core.Dodge:         "Space",
		core.Interact:      "E",
		core.UseItem:       "MouseButtonRight",
		core.UseConsumable: "Q",
		core.Map:           "M",
		core.Menu:          "Escape",
	}
)

type InputHandler struct {
	repo     InputRepository
	log      logs.Logger
	settings InputSettings

	currentKeys   []string
	currentCursor data.Vector
	currentInputs []core.Input
	inputStates   map[core.Input]int
}

type InputState struct {
	Cursor data.Vector
	Keys   []string
}

type InputSettings struct {
	Keyboard map[core.Input]string
}

type InputRepository interface {
	GetInputSettings() (InputSettings, error)
	SetInputSettings(InputSettings) error
}

func NewInputHandler(l logs.Logger, r InputRepository) (*InputHandler, error) {
	if l == nil {
		return nil, ErrLoggerNil
	}

	if r == nil {
		return nil, ErrInputRepoNil
	}

	h := &InputHandler{
		repo: r,
		log:  l,
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

func (h *InputHandler) GetCursor() data.Vector {
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
