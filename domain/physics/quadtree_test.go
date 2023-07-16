package physics_test

import (
	"reflect"
	"testing"

	"github.com/maladroitthief/entree/domain/physics"
)

func TestQuadTree_Clear(t *testing.T) {
	tests := []struct {
		name string
		qt   *physics.QuadTree[int]
	}{
		{
			name: "default",
			qt:   physics.NewQuadTree[int](0, physics.NewRectangle(0, 0, 10, 10)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.qt.Clear()
		})
	}
}

func TestQuadTree_Insert(t *testing.T) {
	type args struct {
		qti *physics.QuadTreeItem[int]
	}
	tests := []struct {
		name string
		qt   *physics.QuadTree[int]
		args args
	}{
		{
			name: "default",
			qt:   physics.NewQuadTree[int](0, physics.NewRectangle(0, 0, 10, 10)),
			args: args{
				qti: physics.NewQuadTreeItem[int](0, physics.NewRectangle(0, 0, 1, 1)),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.qt.Insert(tt.args.qti)
		})
	}
}

func TestQuadTree_Get(t *testing.T) {
	type args struct {
		r physics.Rectangle
	}
	tests := []struct {
		name string
		qt   *physics.QuadTree[int]
		qti  []*physics.QuadTreeItem[int]
		args args
		want []int
	}{
		{
			name: "first quadrant",
			qt:   physics.NewQuadTree[int](0, physics.NewRectangle(0, 0, 10, 10)),
			args: args{
				r: physics.NewRectangle(0, 0, 1, 1),
			},
			qti: []*physics.QuadTreeItem[int]{
				physics.NewQuadTreeItem[int](1, physics.NewRectangle(2, 2, 4, 4)),
				physics.NewQuadTreeItem[int](2, physics.NewRectangle(6, 2, 8, 4)),
				physics.NewQuadTreeItem[int](3, physics.NewRectangle(2, 6, 4, 8)),
				physics.NewQuadTreeItem[int](4, physics.NewRectangle(6, 6, 8, 8)),
			},
			want: []int{1},
		},
		{
			name: "second quadrant",
			qt:   physics.NewQuadTree[int](0, physics.NewRectangle(0, 0, 10, 10)),
			args: args{
				r: physics.NewRectangle(7, 3, 8, 4),
			},
			qti: []*physics.QuadTreeItem[int]{
				physics.NewQuadTreeItem[int](1, physics.NewRectangle(2, 2, 4, 4)),
				physics.NewQuadTreeItem[int](2, physics.NewRectangle(6, 2, 8, 4)),
				physics.NewQuadTreeItem[int](3, physics.NewRectangle(2, 6, 4, 8)),
				physics.NewQuadTreeItem[int](4, physics.NewRectangle(6, 6, 8, 8)),
			},
			want: []int{2},
		},
		{
			name: "third quadrant",
			qt:   physics.NewQuadTree[int](0, physics.NewRectangle(0, 0, 10, 10)),
			args: args{
				r: physics.NewRectangle(3, 7, 4, 8),
			},
			qti: []*physics.QuadTreeItem[int]{
				physics.NewQuadTreeItem[int](1, physics.NewRectangle(2, 2, 4, 4)),
				physics.NewQuadTreeItem[int](2, physics.NewRectangle(6, 2, 8, 4)),
				physics.NewQuadTreeItem[int](3, physics.NewRectangle(2, 6, 4, 8)),
				physics.NewQuadTreeItem[int](4, physics.NewRectangle(6, 6, 8, 8)),
			},
			want: []int{3},
		},
		{
			name: "fourth quadrant",
			qt:   physics.NewQuadTree[int](0, physics.NewRectangle(0, 0, 10, 10)),
			args: args{
				r: physics.NewRectangle(7, 7, 8, 8),
			},
			qti: []*physics.QuadTreeItem[int]{
				physics.NewQuadTreeItem[int](1, physics.NewRectangle(2, 2, 4, 4)),
				physics.NewQuadTreeItem[int](2, physics.NewRectangle(6, 2, 8, 4)),
				physics.NewQuadTreeItem[int](3, physics.NewRectangle(2, 6, 4, 8)),
				physics.NewQuadTreeItem[int](4, physics.NewRectangle(6, 6, 8, 8)),
			},
			want: []int{4},
		},
		{
			name: "overlap",
			qt:   physics.NewQuadTree[int](0, physics.NewRectangle(0, 0, 10, 10)),
			args: args{
				r: physics.NewRectangle(1, 1, 9, 2),
			},
			qti: []*physics.QuadTreeItem[int]{
				physics.NewQuadTreeItem[int](1, physics.NewRectangle(2, 2, 4, 4)),
				physics.NewQuadTreeItem[int](2, physics.NewRectangle(6, 2, 8, 4)),
				physics.NewQuadTreeItem[int](3, physics.NewRectangle(2, 6, 4, 8)),
				physics.NewQuadTreeItem[int](4, physics.NewRectangle(6, 6, 8, 8)),
			},
			want: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reduce the max items to see the actual results
			tt.qt.SetMaxItems(1)
			for _, qti := range tt.qti {
				tt.qt.Insert(qti)
			}

			got := tt.qt.Get(tt.args.r)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuadTree.Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}
