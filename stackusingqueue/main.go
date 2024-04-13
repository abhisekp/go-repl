package stackusingqueue

import (
	. "github.com/abhisekp/go-repl/queue"
	"github.com/abhisekp/go-repl/stack"
)

type Stack[T comparable] struct {
	_      struct{}
	_queue IQueue[T]
}

func (s *Stack[T]) Push(data T) {
	s._queue.Enqueue(data)
	currSize := s.Size()

	for range currSize - 1 {
		s._queue.Enqueue(s._queue.Peek())
		s._queue.Dequeue()
	}
}

func (s *Stack[T]) Pop() {
	s._queue.Dequeue()
}

func (s *Stack[T]) Peek() T {
	return s._queue.Peek()
}

func (s *Stack[T]) Size() int {
	return s._queue.Size()
}

func (s *Stack[T]) IsEmpty() bool {
	return s._queue.IsEmpty()
}

func NewStack[T comparable]() stack.IStack[T] {
	return &Stack[T]{
		_queue: NewQueue[T](),
	}
}

func Run() {

}
