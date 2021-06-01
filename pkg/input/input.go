package input

type Input interface {
	Update()
}

type input struct {
}

func New() Input {
	return &input{}
}

func (i *input) Update() {
	return
}
