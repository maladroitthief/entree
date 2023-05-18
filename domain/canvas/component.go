package canvas

type Component interface {
	Receive(e Entity, message string, value string)
}

type InputComponent interface {
	Component
	Update(Entity)
}

type PhysicsComponent interface {
	Component
	Update(Entity, *Canvas)
}

type GraphicsComponent interface {
	Component
	Update(Entity)
}
