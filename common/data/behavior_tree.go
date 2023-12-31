package data

type Status int
type Archetype int

const (
	RUNNING Status = iota
	SUCCESS
	FAILURE

	COMPOSITE_NODE Archetype = iota
	DECORATOR_NODE
	LEAF_NODE
)

type BehaviorTree struct {
	Nodes []BehaviorTreeNode
	Root  BehaviorTreeNode
}

type BehaviorTreeNode struct {
	Status    Status
	Archetype Archetype
	Children  []BehaviorTreeNode
}

func NewCompositeNode() BehaviorTreeNode {
	return BehaviorTreeNode{
		Status:    RUNNING,
		Archetype: COMPOSITE_NODE,
		Children:  []BehaviorTreeNode{},
	}
}

func NewDecoratorNode() BehaviorTreeNode {
	return BehaviorTreeNode{
		Status:    RUNNING,
		Archetype: DECORATOR_NODE,
		Children:  make([]BehaviorTreeNode, 1),
	}
}

func NewLeafNode() BehaviorTreeNode {
	return BehaviorTreeNode{
		Status:    RUNNING,
		Archetype: LEAF_NODE,
		Children:  make([]BehaviorTreeNode, 0),
	}
}
