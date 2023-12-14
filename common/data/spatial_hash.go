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
	sb        strings.Builder
	Cells     [][]Cell[T]
	X         int
	Y         int
	ChunkSize float64
}

func NewSpatialHash[T comparable](x, y int, size float64) *SpatialHash[T] {
	cells := make([][]Cell[T], x)
	for i := range cells {
		cells[i] = make([]Cell[T], y)
	}

	return &SpatialHash[T]{
		sb:        strings.Builder{},
		X:         x,
		Y:         y,
		ChunkSize: size,
		Cells:     cells,
	}
}

func (h *SpatialHash[T]) Size() int {
	size := 0

	for x := range h.Cells {
		for y := range h.Cells[x] {
			size += len(h.Cells[x][y].items)
		}
	}

	return size
}

func (h *SpatialHash[T]) Insert(val T, bounds Rectangle) {
	minPoint, maxPoint := bounds.MinPoint(), bounds.MaxPoint()
	xMinIndex, yMinIndex := h.toIndex(minPoint.X, minPoint.Y)
	xMaxIndex, yMaxIndex := h.toIndex(maxPoint.X, maxPoint.Y)

	for x, xn := xMinIndex, xMaxIndex; x <= xn; x++ {
		for y, yn := yMinIndex, yMaxIndex; y <= yn; y++ {
			cell := h.Cells[x][y]
			h.Cells[x][y] = cell.Insert(val)
		}
	}
}

func (h *SpatialHash[T]) Update(val T, oldBounds, newBounds Rectangle) {
	h.Delete(val, oldBounds)
	h.Insert(val, newBounds)
}

func (h *SpatialHash[T]) Delete(val T, bounds Rectangle) {
	minPoint, maxPoint := bounds.MinPoint(), bounds.MaxPoint()
	xMinIndex, yMinIndex := h.toIndex(minPoint.X, minPoint.Y)
	xMaxIndex, yMaxIndex := h.toIndex(maxPoint.X, maxPoint.Y)

	for x, xn := xMinIndex, xMaxIndex; x <= xn; x++ {
		for y, yn := yMinIndex, yMaxIndex; y <= yn; y++ {
			cell := h.Cells[x][y]
			h.Cells[x][y] = cell.Delete(val)
		}
	}
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
	xIndex, yIndex := h.toIndex(x, y)
	cell := h.Cells[xIndex][yIndex]

	return cell.Get()
}

func (h *SpatialHash[T]) SearchNeighbors(x, y float64) []T {
	items := []T{}
	for _, direction := range directions {
		i := x + direction[0]*h.ChunkSize
		j := y + direction[1]*h.ChunkSize

		items = append(items, h.Search(i, j)...)
	}

	return items
}

func (h *SpatialHash[T]) Drop() {
	for i := range h.Cells {
		h.Cells[i] = make([]Cell[T], h.Y)
	}
}

func (h *SpatialHash[T]) toIndex(x, y float64) (xIndex, yIndex int) {
	xIndex = int(math.Round(x / h.ChunkSize))
	yIndex = int(math.Round(y / h.ChunkSize))

	xIndex = max(xIndex, 0)
	xIndex = min(xIndex, h.X-1)
	yIndex = max(yIndex, 0)
	yIndex = min(yIndex, h.Y-1)

	return xIndex, yIndex
}

func (h *SpatialHash[T]) toIndices(x, y float64) [][2]int {
	indices := [][2]int{}

	xIndex := int(math.Round(x / h.ChunkSize))
	yIndex := int(math.Round(y / h.ChunkSize))

	xIndex = max(xIndex, 0)
	xIndex = min(xIndex, h.X-1)
	yIndex = max(yIndex, 0)
	yIndex = min(yIndex, h.Y-1)
	indices = append(indices, [2]int{xIndex, yIndex})

	return indices
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

func (c Cell[T]) Insert(item T) Cell[T] {
	c.items = append(
		c.items,
		CellItem[T]{
			item: item,
		},
	)
	return c
}

func (c Cell[T]) Delete(item T) Cell[T] {
	for i := 0; i < len(c.items); i++ {
		if c.items[i].item != item {
			continue
		}
		c.items[i] = c.items[len(c.items)-1]
		c.items = c.items[:len(c.items)-1]
	}

	return c
}

type CellItem[T comparable] struct {
	item T
}
