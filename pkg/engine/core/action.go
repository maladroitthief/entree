package core

type Action interface {
	Execute(Actor)
}
