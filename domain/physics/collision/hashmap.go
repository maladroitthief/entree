package collision

type Index struct {
	X int
	Y int
}

type Hashmap[T comparable] struct {
	Cells      map[Index]*Cell[T]
	ChunkSizeX int
	ChunkSizeY int
}

func NewHashmap[T comparable](sizeX, sizeY int) *Hashmap[T] {
	return &Hashmap[T]{
		ChunkSizeX: sizeX,
		ChunkSizeY: sizeY,
		Cells:      make(map[Index]*Cell[T]),
	}
}

func (h *Hashmap[T]) Clear() {
	for _, c := range h.Cells {
		c.Clear()
	}
}

func (h *Hashmap[T]) GetCell(i Index) *Cell[T] {
	cell, ok := h.Cells[i]
	if !ok {
		cell = NewCell[T]()
		h.Cells[i] = cell
	}
	return cell
}

func (h *Hashmap[T]) ToIndex(x, y int) Index {
	xPos := (x + (h.ChunkSizeX / 2)) / h.ChunkSizeX
	yPos := (y + (h.ChunkSizeY / 2)) / h.ChunkSizeY

	return Index{xPos, yPos}
}

type Cell[T comparable] struct {
	items []T
}

func NewCell[T comparable]() *Cell[T] {
	return &Cell[T]{
		items: make([]T, 0),
	}
}

func (c *Cell[T]) Add(item T) {
	c.items = append(c.items, item)
}

func (c *Cell[T]) Clear() {
	c.items = c.items[:0]
}
