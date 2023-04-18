package canvas_test

import (
	"reflect"
	"testing"

	"github.com/maladroitthief/entree/domain/canvas"
)

func TestNewCanvas(t *testing.T) {
	type args struct {
		x    int
		y    int
		size int
	}
	tests := []struct {
		name string
		args args
		want *canvas.Canvas
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canvas.NewCanvas(tt.args.x, tt.args.y, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCanvas() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCanvas_AddEntity(t *testing.T) {
	type args struct {
		e *canvas.Entity
	}
	tests := []struct {
		name string
		c    *canvas.Canvas
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.AddEntity(tt.args.e)
		})
	}
}

func TestCanvas_Entities(t *testing.T) {
	tests := []struct {
		name string
		c    *canvas.Canvas
		want []*canvas.Entity
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Entities(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Canvas.Entities() = %v, want %v", got, tt.want)
			}
		})
	}
}
