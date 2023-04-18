package sprite_test

import (
	"image"
	"reflect"
	"testing"

	"github.com/maladroitthief/entree/domain/sprite"
)

func TestNewSpriteSheet(t *testing.T) {
	type args struct {
		name    string
		image   image.Image
		rows    int
		columns int
		offset  int
		size    int
	}
	tests := []struct {
		name    string
		args    args
		want    sprite.SpriteSheet
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sprite.NewSpriteSheet(tt.args.name, tt.args.image, tt.args.rows, tt.args.columns, tt.args.offset, tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSpriteSheet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSpriteSheet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_spriteSheet_AddSprite(t *testing.T) {
	type fields struct {
		Name       string
		Image      image.Image
		Rows       int
		Columns    int
		Offset     int
		SpriteSize int
		Sprites    map[string]sprite.Sprite
	}
	type args struct {
		s sprite.Sprite
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss, err := sprite.NewSpriteSheet(
				tt.fields.Name,
				tt.fields.Image,
				tt.fields.Rows,
				tt.fields.Columns,
				tt.fields.Offset,
				tt.fields.SpriteSize,
			)
			if err != nil {
				t.Errorf("NewSpriteSheet() error = %v", err)
				return
			}
			if err := ss.AddSprite(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("spriteSheet.AddSprite() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_spriteSheet_GetName(t *testing.T) {
	type fields struct {
		Name       string
		Image      image.Image
		Rows       int
		Columns    int
		Offset     int
		SpriteSize int
		Sprites    map[string]sprite.Sprite
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss, err := sprite.NewSpriteSheet(
				tt.fields.Name,
				tt.fields.Image,
				tt.fields.Rows,
				tt.fields.Columns,
				tt.fields.Offset,
				tt.fields.SpriteSize,
			)
			if err != nil {
				t.Errorf("NewSpriteSheet() error = %v", err)
				return
			}
			if got := ss.GetName(); got != tt.want {
				t.Errorf("spriteSheet.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_spriteSheet_GetImage(t *testing.T) {
	type fields struct {
		Name       string
		Image      image.Image
		Rows       int
		Columns    int
		Offset     int
		SpriteSize int
		Sprites    map[string]sprite.Sprite
	}
	tests := []struct {
		name   string
		fields fields
		want   image.Image
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss, err := sprite.NewSpriteSheet(
				tt.fields.Name,
				tt.fields.Image,
				tt.fields.Rows,
				tt.fields.Columns,
				tt.fields.Offset,
				tt.fields.SpriteSize,
			)
			if err != nil {
				t.Errorf("NewSpriteSheet() error = %v", err)
				return
			}
			if got := ss.GetImage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spriteSheet.GetImage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_spriteSheet_SpriteRectangle(t *testing.T) {
	type fields struct {
		Name       string
		Image      image.Image
		Rows       int
		Columns    int
		Offset     int
		SpriteSize int
		Sprites    map[string]sprite.Sprite
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    image.Rectangle
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss, err := sprite.NewSpriteSheet(
				tt.fields.Name,
				tt.fields.Image,
				tt.fields.Rows,
				tt.fields.Columns,
				tt.fields.Offset,
				tt.fields.SpriteSize,
			)
			if err != nil {
				t.Errorf("NewSpriteSheet() error = %v", err)
				return
			}
			got, err := ss.SpriteRectangle(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("spriteSheet.SpriteRectangle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spriteSheet.SpriteRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
