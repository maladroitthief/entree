package canvas

type InputComponent interface {
	Update(*Entity)
}

type PhysicsComponent interface {
	Update(*Entity, *Canvas)
}

type GraphicsComponent interface {
	Update(*Entity)
}
