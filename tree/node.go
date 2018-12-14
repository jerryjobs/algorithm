package tree

// IntNodeKey of node
type IntNodeKey int

const (
	left  = 1
	right = -1
)

// Node of tree
type Node struct {
	Key   IntNodeKey
	Value interface{}
}

// Tree interface
type Tree interface {
	Put(key IntNodeKey)
	Get(key IntNodeKey) *Node
}
