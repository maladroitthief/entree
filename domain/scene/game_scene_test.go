package scene_test

import (
	"image/color"
	"reflect"
	"testing"

	"github.com/maladroitthief/entree/domain/canvas"
	"github.com/maladroitthief/entree/domain/scene"
)

func TestNewGameScene(t *testing.T) {
	type args struct {
		state *scene.GameState
	}
	tests := []struct {
		name string
		args args
		want *scene.GameScene
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scene.NewGameScene(tt.args.state); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGameScene() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGameScene_Update(t *testing.T) {
	type args struct {
		state *scene.GameState
	}
	tests := []struct {
		name    string
		s       *scene.GameScene
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Update(tt.args.state); (err != nil) != tt.wantErr {
				t.Errorf("GameScene.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGameScene_GetEntities(t *testing.T) {
	tests := []struct {
		name string
		s    *scene.GameScene
		want []*canvas.Entity
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetEntities(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GameScene.GetEntities() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGameScene_GetBackgroundColor(t *testing.T) {
	tests := []struct {
		name string
		s    *scene.GameScene
		want color.Color
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetBackgroundColor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GameScene.GetBackgroundColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
