package stackusinglinkedlist

import (
	"cmp"
	"fmt"
	"github.com/abhisekp/go-repl/linkedlist"
	"github.com/abhisekp/go-repl/stack"
	"strings"
)

type Stack[T cmp.Ordered] struct {
	_    struct{}
	list linkedlist.ILinkedList[T]
}

func NewStack[T cmp.Ordered]() stack.IStack[T] {
	return &Stack[T]{
		list: linkedlist.NewLinkedList[T](),
	}
}

func (s *Stack[T]) Push(data T) {
	s.list.InsertAtHead(data)
}

func (s *Stack[T]) Pop() {
	s.list.RemoveAtHead()
}

func (s *Stack[T]) Peek() T {
	var zero T
	if s.list.IsEmpty() {
		return zero
	}

	return s.list.Head().Data
}

func (s *Stack[T]) Size() int {
	return s.list.Size()
}

func (s *Stack[T]) IsEmpty() bool {
	return s.list.IsEmpty()
}

func (s *Stack[T]) Print() {
	var str strings.Builder
	str.WriteString("> ")

	list := NewStack[T]()
	for !s.IsEmpty() {
		val := s.Peek()
		list.Push(val)
		s.Pop()

		str.WriteString(fmt.Sprintf("%v -> ", val))
	}
	str.WriteString("|")

	// Re-add everything
	for !list.IsEmpty() {
		s.Push(list.Peek())
		list.Pop()
	}

	fmt.Println(str.String())
}
