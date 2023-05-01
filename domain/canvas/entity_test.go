package canvas_test

import (
	"reflect"
	"testing"

	"github.com/maladroitthief/entree/domain/canvas"
)

type InputMock struct{}

func (m *InputMock) Update(*canvas.Entity)                  {}
func (m *InputMock) Receive(*canvas.Entity, string, string) {}

type PhysicsMock struct{}

func (m *PhysicsMock) Update(*canvas.Entity, *canvas.Canvas)  {}
func (m *PhysicsMock) Receive(*canvas.Entity, string, string) {}

type GraphicsMock struct{}

func (m *GraphicsMock) Update(*canvas.Entity)                  {}
func (m *GraphicsMock) Receive(*canvas.Entity, string, string) {}

func TestEntity_Update(t *testing.T) {
	type args struct {
		c *canvas.Canvas
	}
	tests := []struct {
		name string
		e    *canvas.Entity
		args args
	}{
		{
			name: "default",
			e: &canvas.Entity{
				Input:    &InputMock{},
				Physics:  &PhysicsMock{},
				Graphics: &GraphicsMock{},
			},
			args: args{&canvas.Canvas{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.e.Update(tt.args.c)
		})
	}
}

func TestEntity_Reset(t *testing.T) {
	type want struct {
		state        string
		orientationX canvas.OrientationX
	}
	tests := []struct {
		name string
		e    *canvas.Entity
		want want
	}{
		{
			name: "default",
			e: &canvas.Entity{
				State:        "move",
				OrientationX: canvas.West,
			},
			want: want{
				state:        "idle",
				orientationX: canvas.West,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.e.Reset()
			if !reflect.DeepEqual(tt.e.State, tt.want.state) {
				t.Errorf("Entity.Reset() state = %v, want %v", tt.e.State, tt.want.state)
			}

			if !reflect.DeepEqual(tt.e.OrientationX, tt.want.orientationX) {
				t.Errorf("Entity.Reset() OrientationX = %v, want %v", tt.e.OrientationX, tt.want.orientationX)
			}
		})
	}
}
