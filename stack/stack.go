package stack

import (
	"fmt"
	"strings"
)

type IStack[T any] interface {
	Push(data T)
	Pop()
	Peek() T
	Size() int
	IsEmpty() bool
	Print()
}

func NewStack[T any]() IStack[T] {
	return &Stack[T]{}
}

type Stack[T any] struct {
	_      struct{}
	_stack []T
}

func (s *Stack[T]) Pop() {
	if s.Size() == 0 {
		return
	}

	s._stack = s._stack[:s.Size()-1]
}

func (s *Stack[T]) Peek() T {
	var empty T

	if s.Size() == 0 {
		return empty
	}

	return s._stack[s.Size()-1]
}

func (s *Stack[T]) Size() int {
	return len(s._stack)
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack[T]) Print() {
	var str strings.Builder
	str.WriteString("> ")
	for _, data := range s._stack {
		str.WriteString(fmt.Sprintf("%v -> ", data))
	}
	str.WriteString("|")

	fmt.Println(str.String())
}

func (s *Stack[T]) Push(data T) {
	s._stack = append(s._stack, data)
}
