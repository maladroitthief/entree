package infrastructure

import (
	"encoding/json"
	"io/ioutil"

	"github.com/maladroitthief/entree/domain/sprite"
)

type SpriteSheetModel struct {
	Name       string        `json:"Name"`
	Path       string        `json:"Path"`
	Rows       int           `json:"Rows"`
	Columns    int           `json:"Columns"`
	SpriteSize int           `json:"SpriteSize"`
	Sprites    []SpriteModel `json:"Sprites"`
}

type SpriteModel struct {
	Name   string `json:"Name"`
	Row    int    `json:"Row"`
	Column int    `json:"Column"`
}

type SpriteJsonRepository struct {
}

func NewSpriteJsonRepository() *SpriteJsonRepository {
	r := &SpriteJsonRepository{}

	return r
}

func (r *SpriteJsonRepository) GetSpriteSheet(path string) (sprite.SpriteSheet, error) {
	jsonContent, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	ss := SpriteSheetModel{}
	err = json.Unmarshal(jsonContent, &ss)
	if err != nil {
		return nil, err
	}

	return r.unmarshalSpriteSheet(ss)
}

func (r *SpriteJsonRepository) unmarshalSpriteSheet(
	ssm SpriteSheetModel,
) (sprite.SpriteSheet, error) {
	ss, err := sprite.NewSpriteSheet(
		ssm.Name,
		ssm.Path,
		ssm.Rows,
		ssm.Columns,
		ssm.SpriteSize,
	)
  if err != nil {
    return nil, err
  }

	for _, sm := range ssm.Sprites {
		s := r.unmarshalSprite(sm)
		err := ss.AddSprite(s)
		if err != nil {
			return nil, err
		}
	}

	return ss, nil
}

func (r *SpriteJsonRepository) unmarshalSprite(
	m SpriteModel,
) sprite.Sprite {
	s := sprite.Sprite{
		Name:   m.Name,
		Row:    m.Row,
		Column: m.Column,
	}

	return s
}
