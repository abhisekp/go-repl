package stack

type IStack[T any] interface {
	Push(data T)
	Pop()
	Peek() T
	Size() int
	IsEmpty() bool
}
