package data

import (
	"math"
	"strings"
)

var (
	directions = [][2]float64{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 0},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
)

type SpatialHash[T comparable] struct {
	sb         strings.Builder
	Cells      map[[2]float64]Cell[T]
	ChunkSizeX float64
	ChunkSizeY float64
}

func NewSpatialHash[T comparable](sizeX, sizeY float64) *SpatialHash[T] {
	return &SpatialHash[T]{
		sb:         strings.Builder{},
		ChunkSizeX: sizeX,
		ChunkSizeY: sizeY,
		Cells:      make(map[[2]float64]Cell[T]),
	}
}

func (h *SpatialHash[T]) Insert(val T, position Vector) {
	positionIndex := h.toIndex(position.X, position.Y)
	cell, ok := h.Cells[positionIndex]

	if !ok {
		cell = NewCell[T]()
	}

	h.Cells[positionIndex] = cell.Insert(val, position)
}

func (h *SpatialHash[T]) Update(val T, oldPosition, newPosition Vector) {
	h.Delete(val, oldPosition)
	h.Insert(val, newPosition)
}

func (h *SpatialHash[T]) Delete(val T, position Vector) {
	positionIndex := h.toIndex(position.X, position.Y)
	cell, ok := h.Cells[positionIndex]

	if !ok {
		cell = NewCell[T]()
	}

	h.Cells[positionIndex] = cell.Delete(val, position)
}

func (h *SpatialHash[T]) WalkGrid(v, w Vector) []Vector {
	delta := w.Subtract(v)
	nX, nY := math.Abs(delta.X), math.Abs(delta.Y)
	signX, signY := 1.0, 1.0
	if delta.X <= 0 {
		signX = -1
	}
	if delta.Y <= 0 {
		signY = -1
	}
	vector := v.Clone()
	vectors := []Vector{vector.Clone()}

	i, j := 0.0, 0.0
	for i < nX || j < nY {
		if (1+2*i)*nY < (1+2*j)*nX {
			vector.X += signX
			i++
		} else {
			vector.Y += signY
			j++
		}
		vectors = append(vectors, vector.Clone())
	}

	return vectors
}

func (h *SpatialHash[T]) Search(x, y float64) []T {
	i := h.toIndex(x, y)

	cell, ok := h.Cells[i]
	if !ok {
		cell = NewCell[T]()
		h.Cells[i] = cell
	}

	return cell.Get()
}

func (h *SpatialHash[T]) SearchNeighbors(x, y float64) []T {
	items := []T{}
	for _, direction := range directions {
		i := x + direction[0]*h.ChunkSizeX
		j := y + direction[1]*h.ChunkSizeY

		items = append(items, h.Search(i, j)...)
	}

	return items
}

func (h *SpatialHash[T]) Drop() {
	for k, c := range h.Cells {
		h.Cells[k] = c.Drop()
	}
}

func (h *SpatialHash[T]) toIndex(x, y float64) [2]float64 {
	xIndex := math.Round(x/h.ChunkSizeX) * h.ChunkSizeX
	yIndex := math.Round(y/h.ChunkSizeY) * h.ChunkSizeY

	return [2]float64{xIndex, yIndex}
}

type Cell[T comparable] struct {
	items []CellItem[T]
}

func NewCell[T comparable]() Cell[T] {
	return Cell[T]{
		items: make([]CellItem[T], 0),
	}
}

func (c Cell[T]) Get() []T {
	items := make([]T, len(c.items))

	for i := 0; i < len(c.items); i++ {
		items[i] = c.items[i].item
	}

	return items
}

func (c Cell[T]) Insert(item T, position Vector) Cell[T] {
	c.items = append(
		c.items,
		CellItem[T]{
			item:     item,
			position: position,
		},
	)

	return c
}

func (c Cell[T]) Delete(item T, position Vector) Cell[T] {
	for i := 0; i < len(c.items); i++ {
		if c.items[i].item != item {
			continue
		}

		c.items[i] = c.items[len(c.items)-1]
		c.items = c.items[:len(c.items)-1]
	}

	return c
}

func (c Cell[T]) Drop() Cell[T] {
	c.items = c.items[:0]
	return c
}

type CellItem[T comparable] struct {
	position Vector
	item     T
}
