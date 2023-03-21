package sprite

type Repository interface {
	GetSpriteSheet(path string) (SpriteSheet, error)
}
