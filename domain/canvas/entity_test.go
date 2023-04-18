package canvas_test

import (
	"reflect"
	"testing"

	"github.com/maladroitthief/entree/domain/canvas"
)

type InputMock struct {
}

func (m *InputMock) Update(*canvas.Entity) {
}

type PhysicsMock struct {
}

func (m *PhysicsMock) Update(*canvas.Entity, *canvas.Canvas) {
}

type GraphicsMock struct {
}

func (m *GraphicsMock) Update(*canvas.Entity) {
}

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

func TestEntity_VariantUpdate(t *testing.T) {
	tests := []struct {
		name string
		e    *canvas.Entity
		want int
	}{
		{
			name: "default",
			e: &canvas.Entity{
				SpriteSpeed:       1,
				StateCounter:      1,
				SpriteMaxVariants: 1,
				SpriteVariant:     1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.e.VariantUpdate()
			got := tt.e.SpriteVariant

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.VariantUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_Reset(t *testing.T) {
	type want struct {
		state        string
		orientationX canvas.OrientationX
		deltaX       float64
		deltaY       float64
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
				DeltaX:       1,
				DeltaY:       1,
			},
			want: want{
				state:        "idle",
				orientationX: canvas.West,
				deltaX:       0,
				deltaY:       0,
			},
		},
		{
			name: "no deltaX movement",
			e: &canvas.Entity{
				State:        "move",
				OrientationX: canvas.West,
				DeltaX:       0,
				DeltaY:       1,
			},
			want: want{
				state:        "idle",
				orientationX: canvas.Neutral,
				deltaX:       0,
				deltaY:       0,
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

			if !reflect.DeepEqual(tt.e.DeltaX, tt.want.deltaX) {
				t.Errorf("Entity.Reset() DeltaX = %v, want %v", tt.e.DeltaX, tt.want.deltaX)
			}

			if !reflect.DeepEqual(tt.e.DeltaY, tt.want.deltaY) {
				t.Errorf("Entity.Reset() DeltaY = %v, want %v", tt.e.DeltaY, tt.want.deltaY)
			}
		})
	}
}
