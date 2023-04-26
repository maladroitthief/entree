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

func (h *Hashmap[T]) Add(val T, bounds Rectangle) {
	min := h.ToIndex(int(bounds.MinPoint.X), int(bounds.MinPoint.Y))
	max := h.ToIndex(int(bounds.MaxPoint.X), int(bounds.MaxPoint.Y))

	for x := min.X; x <= max.X; x++ {
		for y := min.Y; y <= max.Y; y++ {
			cell := h.GetCell(Index{x, y})
			cell.Add(val, bounds)
		}
	}
}

func (h *Hashmap[T]) Check(bounds Rectangle) []T {
	min := h.ToIndex(int(bounds.MinPoint.X), int(bounds.MinPoint.Y))
	max := h.ToIndex(int(bounds.MaxPoint.X), int(bounds.MaxPoint.Y))

	flags := make(map[T]bool)

	for x := min.X; x <= max.X; x++ {
		for y := min.Y; y <= max.Y; y++ {
			cell := h.GetCell(Index{x, y})
			collisions := cell.Check(bounds)
			for _, collision := range collisions {
				flags[collision.item] = true
			}
		}
	}

	items := make([]T, 0, len(flags))
	for item := range flags {
		items = append(items, item)
	}

	return items
}

type Cell[T comparable] struct {
	items []CellItem[T]
}

func NewCell[T comparable]() *Cell[T] {
	return &Cell[T]{
		items: make([]CellItem[T], 0),
	}
}

func (c *Cell[T]) Add(item T, bounds Rectangle) {
	c.items = append(
		c.items,
		CellItem[T]{
			item:   item,
			bounds: bounds,
		},
	)
}

func (c *Cell[T]) Check(bounds Rectangle) []CellItem[T] {
	items := make([]CellItem[T], 0)
	for _, item := range c.items {
		if bounds.Intersects(item.bounds) {
			items = append(items, item)
		}
	}

	return items
}

func (c *Cell[T]) Clear() {
	c.items = c.items[:0]
}

type CellItem[T comparable] struct {
	bounds Rectangle
	item   T
}
