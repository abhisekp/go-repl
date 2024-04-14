package stackusinglinkedlist

import (
	"github.com/abhisekp/go-repl/linkedlist"
	"github.com/abhisekp/go-repl/stack"
)

type Stack[T comparable] struct {
	_     struct{}
	_list linkedlist.ILinkedList[T]
}

func NewStack[T comparable]() stack.IStack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(data T) {
	// TODO implement me
	panic("implement me")
}

func (s *Stack[T]) Pop() {
	// TODO implement me
	panic("implement me")
}

func (s *Stack[T]) Peek() T {
	// TODO implement me
	panic("implement me")
}

func (s *Stack[T]) Size() int {
	return s._list.Size()
}

func (s *Stack[T]) IsEmpty() bool {
	// TODO implement me
	panic("implement me")
}

func (s *Stack[T]) Print() {
	// TODO implement me
	panic("implement me")
}
