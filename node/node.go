package node

type Node[T any] struct {
	Data T
	Next *Node[T]
}

type DLLNode[T any] struct {
	Data T
	Prev *DLLNode[T]
	Next *DLLNode[T]
}
