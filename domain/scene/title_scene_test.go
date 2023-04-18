package scene_test

import (
	"image/color"
	"reflect"
	"testing"

	"github.com/maladroitthief/entree/domain/canvas"
	"github.com/maladroitthief/entree/domain/scene"
)

func TestNewTitleScene(t *testing.T) {
	type args struct {
		state *scene.GameState
	}
	tests := []struct {
		name string
		args args
		want *scene.TitleScene
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scene.NewTitleScene(tt.args.state); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTitleScene() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTitleScene_Update(t *testing.T) {
	type args struct {
		state *scene.GameState
	}
	tests := []struct {
		name    string
		s       *scene.TitleScene
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Update(tt.args.state); (err != nil) != tt.wantErr {
				t.Errorf("TitleScene.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTitleScene_GetEntities(t *testing.T) {
	tests := []struct {
		name string
		s    *scene.TitleScene
		want []*canvas.Entity
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetEntities(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TitleScene.GetEntities() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTitleScene_GetBackgroundColor(t *testing.T) {
	tests := []struct {
		name string
		s    *scene.TitleScene
		want color.Color
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetBackgroundColor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TitleScene.GetBackgroundColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
