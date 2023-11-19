package data

const (
	DefaultMaxItems  = 500
	DefaultMaxLevels = 10
)

type QuadTree[T comparable] struct {
	maxItems  int
	maxLevels int
	level     int
	items     []*QuadTreeItem[T]
	bounds    Rectangle
	quadrants [4]*QuadTree[T]
}

type QuadTreeItem[T comparable] struct {
	bounds Rectangle
	item   T
}

func NewQuadTree[T comparable](level int, bounds Rectangle) *QuadTree[T] {
	qt := &QuadTree[T]{
		maxItems:  DefaultMaxItems,
		maxLevels: DefaultMaxLevels,
		level:     level,
		bounds:    bounds,
		items:     make([]*QuadTreeItem[T], 0),
	}

	return qt
}

func NewQuadTreeItem[T comparable](item T, bounds Rectangle) *QuadTreeItem[T] {
	return &QuadTreeItem[T]{
		item:   item,
		bounds: bounds,
	}
}

func (q *QuadTree[T]) SetMaxItems(max int) {
	q.maxItems = max
}

func (q *QuadTree[T]) SetMaxLevels(max int) {
	q.maxLevels = max
}

func (q *QuadTree[T]) Clear() {
	q.items = q.items[:0]

	for _, quadrant := range q.quadrants {
		if quadrant != nil {
			quadrant.Clear()
		}
		quadrant = nil
	}
}

func (q *QuadTree[T]) split() {
	q.quadrants[0] = NewQuadTree[T](
		q.level+1,
		NewRectangle(
			Vector{X: q.bounds.Position.X - q.halfWidth(), Y: q.bounds.Position.Y - q.halfHeight()},
			q.halfWidth(),
			q.halfHeight(),
		),
	)
	q.quadrants[1] = NewQuadTree[T](
		q.level+1,
		NewRectangle(
			Vector{X: q.bounds.Position.X - q.halfWidth(), Y: q.bounds.Position.Y + q.halfHeight()},
			q.halfWidth(),
			q.halfHeight(),
		),
	)
	q.quadrants[2] = NewQuadTree[T](
		q.level+1,
		NewRectangle(
			Vector{X: q.bounds.Position.X + q.halfWidth(), Y: q.bounds.Position.Y + q.halfHeight()},
			q.halfWidth(),
			q.halfHeight(),
		),
	)
	q.quadrants[3] = NewQuadTree[T](
		q.level+1,
		NewRectangle(
			Vector{X: q.bounds.Position.X + q.halfWidth(), Y: q.bounds.Position.Y - q.halfHeight()},
			q.halfWidth(),
			q.halfHeight(),
		),
	)
}

func (q *QuadTree[T]) Index(r Rectangle) int {
	if q.quadrants[0] == nil {
		return -1
	}

	if q.quadrants[0].bounds.Intersects(r) {
		return 0
	}

	if q.quadrants[1].bounds.Intersects(r) {
		return 1
	}

	if q.quadrants[2].bounds.Intersects(r) {
		return 2
	}

	if q.quadrants[3].bounds.Intersects(r) {
		return 3
	}

	return -1
}

func (q *QuadTree[T]) InclusiveIndexes(r Rectangle) []int {
	indexes := []int{}

	if q.quadrants[0] == nil {
		return append(indexes, -1)
	}

	if q.quadrants[0].bounds.Intersects(r) {
		indexes = append(indexes, 0)
	}

	if q.quadrants[1].bounds.Intersects(r) {
		indexes = append(indexes, 1)
	}

	if q.quadrants[2].bounds.Intersects(r) {
		indexes = append(indexes, 2)
	}

	if q.quadrants[3].bounds.Intersects(r) {
		indexes = append(indexes, 3)
	}

	if len(indexes) == 0 {
		indexes = append(indexes, -1)
	}

	return indexes
}

func (q *QuadTree[T]) Insert(item T, bounds Rectangle) {
	qi := NewQuadTreeItem[T](item, bounds)

	if q.quadrants[0] != nil {
		index := q.Index(qi.bounds)
		if index != -1 {
			q.quadrants[index].Insert(qi.item, qi.bounds)
			return
		}
	}
	q.items = append(q.items, qi)

	if len(q.items) <= q.maxItems || q.level >= q.maxLevels {
		return
	}

	if q.quadrants[0] == nil {
		q.split()
	}

	i := 0
	for i < len(q.items) {
		index := q.Index(q.items[i].bounds)
		if index != -1 {
			q.quadrants[index].Insert(q.items[i].item, q.items[i].bounds)
			copy(q.items[i:], q.items[i+1:])
			q.items[len(q.items)-1] = nil
			q.items = q.items[:len(q.items)-1]
		} else {
			i++
		}
	}
}

func (q *QuadTree[T]) Get(r Rectangle) []T {
	results := make([]T, 0)
	indexes := q.InclusiveIndexes(r)

	for _, index := range indexes {
		if index != -1 && q.quadrants[0] != nil {
			results = append(results, q.quadrants[index].Get(r)...)
		}

		for _, i := range q.items {
			results = append(results, i.item)
		}
	}

	return results
}

func (q *QuadTree[T]) halfHeight() float64 {
	return q.bounds.Height / 2
}

func (q *QuadTree[T]) halfWidth() float64 {
	return q.bounds.Width / 2
}
