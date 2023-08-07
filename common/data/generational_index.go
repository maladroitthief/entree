package data

type GenerationalIndex struct {
	index      int
	generation int
}

type allocatorEntry struct {
	isLive     bool
	generation int
}

type GenerationalIndexAllocator struct {
	entries []allocatorEntry
	free    []int
}

func NewGenerationalIndexAllocator() *GenerationalIndexAllocator {
	return &GenerationalIndexAllocator{
		entries: []allocatorEntry{},
		free:    []int{},
	}
}

func (g *GenerationalIndexAllocator) Allocate() GenerationalIndex {
	if len(g.free) <= 0 {
		g.entries = append(g.entries, allocatorEntry{isLive: true, generation: 1})

		return GenerationalIndex{
			index:      len(g.entries) - 1,
			generation: 1,
		}
	}

	n := len(g.free) - 1
	index := g.free[n]
	g.free = g.free[:n]
	g.entries[index].generation++
	g.entries[index].isLive = true

	return GenerationalIndex{
		index:      index,
		generation: g.entries[index].generation,
	}
}

func (g *GenerationalIndexAllocator) Deallocate(i GenerationalIndex) bool {
	if g.IsLive(i) == true {
		g.entries[i.index].isLive = false
		g.free = append(g.free, i.index)

		return true
	}

	return false
}

func (g *GenerationalIndexAllocator) IsLive(i GenerationalIndex) bool {
	if i.index >= len(g.entries) {
		return false
	}

	if g.entries[i.index].generation != i.generation {
		return false
	}

	if g.entries[i.index].isLive == false {
		return false
	}

	return true
}

type ArrayEntry[T comparable] struct {
	value      T
	generation int
}

type GenerationalIndexArray[T comparable] []ArrayEntry[T]

func NewGenerationalIndexArray[T comparable]() GenerationalIndexArray[T] {
	return GenerationalIndexArray[T]{}
}

func (g GenerationalIndexArray[T]) Set(
	index GenerationalIndex,
	value T,
) GenerationalIndexArray[T] {

	for len(g) <= index.index {
		g = append(g, ArrayEntry[T]{generation: -1})
	}

	g[index.index] = ArrayEntry[T]{
		value:      value,
		generation: index.generation,
	}

	return g
}

func (g GenerationalIndexArray[T]) Remove(
	index GenerationalIndex,
) GenerationalIndexArray[T] {

	if index.index < len(g) {
		g[index.index].generation = -1
	}

	return g
}

func (g GenerationalIndexArray[T]) Get(index GenerationalIndex) T {
	var defaultValue T
	if index.index >= len(g) {
		return defaultValue
	}

	entry := g[index.index]
	if entry.generation != index.generation {
		return defaultValue
	}

	return entry.value
}

func (g GenerationalIndexArray[T]) GetAllIndices(
	a *GenerationalIndexAllocator,
) []GenerationalIndex {

	result := []GenerationalIndex{}

	for i, entry := range g {
		if entry.generation <= 0 {
			continue
		}

		index := GenerationalIndex{index: i, generation: entry.generation}
		if a.IsLive(index) {
			result = append(result, index)
		}
	}

	return result
}

func (g GenerationalIndexArray[T]) GetAll(
	a *GenerationalIndexAllocator,
) []T {
	result := []T{}

	for i, entry := range g {
		if entry.generation <= 0 {
			continue
		}

		index := GenerationalIndex{index: i, generation: entry.generation}
		if a.IsLive(index) {
			result = append(result, entry.value)
		}
	}

	return result
}

func (g GenerationalIndexArray[T]) First(
	a *GenerationalIndexAllocator,
) (GenerationalIndex, T) {

	var defaultValue T
	for i, entry := range g {
		if entry.generation <= 0 {
			continue
		}

		index := GenerationalIndex{index: i, generation: entry.generation}
		if a.IsLive(index) {
			return index, entry.value
		}
	}

	return GenerationalIndex{}, defaultValue
}
