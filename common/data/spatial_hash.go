package data

import (
	"fmt"
	"math"
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
	Cells      map[string]Cell[T]
	ChunkSizeX float64
	ChunkSizeY float64
}

func NewSpatialHash[T comparable](sizeX, sizeY float64) *SpatialHash[T] {
	return &SpatialHash[T]{
		ChunkSizeX: sizeX,
		ChunkSizeY: sizeY,
		Cells:      make(map[string]Cell[T]),
	}
}

func (h *SpatialHash[T]) Insert(val T, bounds Rectangle) {
	center := bounds.Center()
	centerIndex := h.toIndex(center.X, center.Y)
	cell, ok := h.Cells[centerIndex]

	if !ok {
		cell = NewCell[T]()
	}

	h.Cells[centerIndex] = cell.Insert(val, bounds)
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

func (h *SpatialHash[T]) toChunk(x, y float64) (xPos, yPos float64) {
	xPos = math.Round(x/h.ChunkSizeX) * h.ChunkSizeX
	yPos = math.Round(y/h.ChunkSizeY) * h.ChunkSizeY

	return xPos, yPos
}

func (h *SpatialHash[T]) toIndex(x, y float64) string {
	xPos, yPos := h.toChunk(x, y)

	return fmt.Sprintf("%v,%v", xPos, yPos)
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
	items := []T{}
	for _, item := range c.items {
		items = append(items, item.item)
	}

	return items
}

func (c Cell[T]) Insert(item T, bounds Rectangle) Cell[T] {
	c.items = append(
		c.items,
		CellItem[T]{
			item:   item,
			bounds: bounds,
		},
	)

	return c
}

func (c Cell[T]) Drop() Cell[T] {
	c.items = c.items[:0]
	return c
}

type CellItem[T comparable] struct {
	bounds Rectangle
	item   T
}
