package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/maladroitthief/entree/pkg/input"
	"github.com/maladroitthief/entree/pkg/scene"
)

func init() {

}

type Game interface {
	Draw(screen *ebiten.Image)
	Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int)
	Update() error
}

// Game holds together all the necessary pieces to keep this show rolling
type game struct {
	input input.Input
	sh    scene.Handler
}

// New constructs our game object
func New() (Game, error) {
	g := game{
		input: input.New(),
		sh:    scene.NewHandler(),
	}

	g.sh.GoTo(&scene.TitleScene{})

	return &g, nil
}

// Draw wraps the scene managers draw function
func (g *game) Draw(screen *ebiten.Image) {
	g.sh.Draw(screen)
}

// Layout takes the window size and returns the screen size
func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

// Update handles all game updates. Works as our game loop
func (g *game) Update() error {
	// Update input
	g.input.Update()

	// Update scene manager
	err := g.sh.Update(g.input)
	if err != nil {
		return err
	}

	return nil
}
