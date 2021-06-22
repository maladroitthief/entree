package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	transitionMaxCount = 20
)

var (
	transitionFrom *ebiten.Image
	transitionTo   *ebiten.Image
)

type Handler interface {
	SetScreenSize(w, h int)
	Draw(r *ebiten.Image)
	GoTo(scene Scene)
	Update(Input) error
}

type handler struct {
	screenWidth     int
	screenHeight    int
	current         Scene
	next            Scene
	transitionCount int
}

type gameState struct {
	sm Handler
	i  Input
}

func NewHandler() Handler {
	return &handler{}
}

func (s *handler) SetScreenSize(w, h int) {
	transitionFrom = ebiten.NewImage(w, h)
	transitionTo = ebiten.NewImage(w, h)

	s.screenWidth = w
	s.screenHeight = h
}

// Update handles the scene changes as well as the input changes
func (h *handler) Update(input Input) error {
	if h.transitionCount == 0 {
		return h.current.Update(&gameState{
			sm: h,
			i:  input,
		})
	}

	h.transitionCount--
	if h.transitionCount > 0 {
		return nil
	}

	h.current = h.next
	h.next = nil
	return nil
}

// Draw places the image in the scene
func (h *handler) Draw(r *ebiten.Image) {
	if h.transitionCount == 0 {
		h.current.Draw(r)
		return
	}

	transitionFrom.Clear()
	h.current.Draw(transitionFrom)

	transitionTo.Clear()
	h.next.Draw(transitionTo)

	r.DrawImage(transitionFrom, nil)

	alpha := 1 - float64(h.transitionCount)/float64(transitionMaxCount)
	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, alpha)
	r.DrawImage(transitionTo, op)
}

// GoTo handles scene changes
func (h *handler) GoTo(scene Scene) {
	if h.current == nil {
		h.current = scene
	} else {
		h.next = scene
		h.transitionCount = transitionMaxCount
	}
}
