package collision

const (
	DefaultMaxItems  = 50
	DefaultMaxLevels = 10
)

type QuadTree[T comparable] struct {
	maxItems  int
	maxLevels int
	level     int
	items     []*QuadTreeItem[T]
	bounds    Rectangle
	nodes     [4]*QuadTree[T]
}

type QuadTreeItem[T comparable] struct {
	bounds Rectangle
	item   T
}

func NewQuadTree[T comparable](level int, bounds Rectangle) *QuadTree[T] {
	return &QuadTree[T]{
		maxItems:  DefaultMaxItems,
		maxLevels: DefaultMaxLevels,
		level:     level,
		bounds:    bounds,
		items:     make([]*QuadTreeItem[T], 0),
	}
}

func (q *QuadTree[T]) Clear() {
	q.items = q.items[:0]

	// TODO: do we need to remove the nodes?
	for _, node := range q.nodes {
		node.Clear()
		node = nil
	}
}

func (q *QuadTree[T]) Split() {
	q.nodes[0] = NewQuadTree[T](
		q.level+1,
		NewRectangle(q.x(), q.y(), q.subWidth(), q.subHeight()),
	)
	q.nodes[1] = NewQuadTree[T](
		q.level+1, NewRectangle(q.x()+q.subWidth(), q.y(), q.subWidth(), q.subHeight()),
	)
	q.nodes[2] = NewQuadTree[T](
		q.level+1,
		NewRectangle(q.x(), q.y()+q.subHeight(), q.subWidth(), q.subHeight()),
	)
	q.nodes[3] = NewQuadTree[T](
		q.level+1,
		NewRectangle(q.x()+q.subWidth(), q.y()+q.subHeight(), q.subWidth(), q.subHeight()),
	)
}

func (q *QuadTree[T]) Index(r Rectangle) int {
	index := -1
	midWidth := q.x() + q.subWidth()
	midHeight := q.y() + q.subHeight()

	inTop := r.MinPoint.Y < midHeight && r.MinPoint.Y+r.Height() < midHeight
	inBottom := r.MinPoint.Y > midHeight
	inLeft := r.MinPoint.X < midWidth && r.MinPoint.X+r.Width() < midWidth
	inRight := r.MinPoint.X > midWidth

	if inLeft && inTop {
		index = 0
	}
	if inRight && inTop {
		index = 1
	}
	if inLeft && inBottom {
		index = 2
	}
	if inRight && inBottom {
		index = 3
	}

	return index
}

func (q *QuadTree[T]) Insert(i *QuadTreeItem[T]) {
	// Attempt to add to child nodes if they exist
	if q.nodes[0] != nil {
		index := q.Index(i.bounds)
		if index != -1 {
			q.nodes[index].Insert(i)
			return
		}
	}

	q.items = append(q.items, i)
	if len(q.items) > q.maxItems && q.level < q.maxLevels {
		if q.nodes[0] == nil {
			q.Split()
		}

		i := 0
		for i < len(q.items) {
			index := q.Index(q.items[i].bounds)
			if index != -1 {
				q.nodes[index].Insert(q.items[i])
				// Remove the item at the given index
				copy(q.items[i:], q.items[i+1:])
				q.items[len(q.items)-1] = nil
				q.items = q.items[:len(q.items)-1]
			} else {
				i++
			}
		}
	}
}

func (q *QuadTree[T]) Get(r Rectangle) []T {
	results := make([]T, 0)
	index := q.Index(r)

	if index != -1 && q.nodes[0] != nil {
		results = q.nodes[index].Get(r)
	}

	for _, i := range q.items {
		results = append(results, i.item)
	}

	return results
}

func (q *QuadTree[T]) x() float64 {
	return q.bounds.MinPoint.X
}

func (q *QuadTree[T]) y() float64 {
	return q.bounds.MinPoint.Y
}

func (q *QuadTree[T]) subHeight() float64 {
	return q.bounds.Height() / 2
}

func (q *QuadTree[T]) subWidth() float64 {
	return q.bounds.Width() / 2
}
