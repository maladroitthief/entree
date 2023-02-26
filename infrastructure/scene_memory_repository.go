package infrastructure

import "github.com/maladroitthief/entree/domain/scene"

type SceneMemoryRepository struct {
	current         *scene.Scene
	next            *scene.Scene
	transitionCount int
}

func NewSceneMemoryRepository() *SceneMemoryRepository {
	return &SceneMemoryRepository{}
}

func (r *SceneMemoryRepository) GetCurrentScene() *scene.Scene {
	return r.current
}

func (r *SceneMemoryRepository) SetCurrentScene(s *scene.Scene) {
	r.current = s
}

func (r *SceneMemoryRepository) GetNextScene() *scene.Scene {
	return r.next
}

func (r *SceneMemoryRepository) SetNextScene(s *scene.Scene) {
	r.next = s
}

func (r *SceneMemoryRepository) GetTransitionCount() int {
	return r.transitionCount
}

func (r *SceneMemoryRepository) SetTransitionCount(t int) {
	r.transitionCount = t
}
