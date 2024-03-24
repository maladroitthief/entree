package data

type (
	Graph[T any] interface {
		Nodes() []Node[T]
		Edges(Node[T]) []Node[T]
	}

	Node[T any] interface {
		Items() []T
	}
)

func BFS[T any](
	graph Graph[T],
	start Node[T],
	maxDepth int,
	process func([]T) error,
) error {
	visited := map[Node[T]]struct{}{}
	queue := NewQueue[Node[T]]()
	queue.Enqueue(start)

	currentDepth := 0
	for queue.Len() > 0 {
		if currentDepth > maxDepth {
			return ErrMaxDepthReached
		}

		nodesAtDepth := queue.Len()
		for i := 0; i < nodesAtDepth; i++ {
			currentNode, err := queue.Dequeue()
			if err != nil {
				return err
			}
			_, ok := visited[currentNode]
			if ok {
				continue
			}
			visited[currentNode] = struct{}{}

			err = process(currentNode.Items())
			if err != nil {
				return err
			}

			edges := graph.Edges(currentNode)
			if len(edges) <= 0 {
				continue
			}

			for _, edge := range edges {
				queue.Enqueue(edge)
			}
		}
		currentDepth++
	}

	return nil
}
