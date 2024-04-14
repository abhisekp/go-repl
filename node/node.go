package node

type Node[T comparable] struct {
	Data T
	Next *Node[T]
}

type DLLNode[T comparable] struct {
	Data T
	Prev *DLLNode[T]
	Next *DLLNode[T]
}
