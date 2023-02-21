package scene

type Scene interface {
	Update() error
	Draw(screen *ebiten.Image)
}

type 
