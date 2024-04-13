package node

type Node[T any] struct {
	Data T
	Next *Node[T]
}

type DLLNode[T any] struct {
	Node[T]
	Prev *Node[T]
}
