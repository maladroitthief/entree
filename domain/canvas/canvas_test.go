package canvas_test

import (
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
	}{
		{
			name: "default",
			args: args{
				x:    8,
				y:    8,
				size: 16,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			canvas.NewCanvas(tt.args.x, tt.args.y, tt.args.size)
		})
	}
}

func TestCanvas_AddEntity(t *testing.T) {
	type args struct {
		e canvas.Entity
	}
	tests := []struct {
		name string
		c    *canvas.Canvas
		args args
	}{
		{
			name: "Default",
			c:    canvas.NewCanvas(0, 0, 0),
			args: args{
				e: &EntityMock{},
			},
		},
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
		len  int
		want []canvas.Entity
	}{
		{
			name: "default",
			c:    canvas.NewCanvas(0, 0, 0),
			len:  0,
			want: []canvas.Entity{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, e := range tt.want {
				tt.c.AddEntity(e)
			}

			got := tt.c.Entities()
			if len(got) != tt.len {
				t.Errorf("Canvas.Entities() = %v, want %v", got, tt.want)
			}
		})
	}
}
