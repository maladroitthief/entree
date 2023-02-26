package scene

type Repository interface {
  GetCurrentScene() *Scene
  SetCurrentScene(*Scene) 
  GetNextScene() *Scene
  SetNextScene(*Scene)
  GetTransitionCount() int
  SetTransitionCount(int)
}
