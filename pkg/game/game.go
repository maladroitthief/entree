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

// Draw passes the rendering responsibility to the scene manager
func (g *game) Draw(screen *ebiten.Image) {
	g.sh.Draw(screen)
}

// Layout takes the window size and returns the screen size
func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

// Update handles the input updates as well as the scene manager updates
func (g *game) Update() error {
	// update input
	g.input.Update()

	// update sceneManager
	err := g.sh.Update(g.input)
	if err != nil {
		return err
	}

	return nil
}
