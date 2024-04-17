package queueusingstack

import (
	"github.com/abhisekp/go-repl/queue"
	. "github.com/abhisekp/go-repl/stack"
)

var _ queue.IQueue[int] = &Queue[int]{}

type Queue[T comparable] struct {
	_      struct{}
	stack1 IStack[T]
	stack2 IStack[T]
}

func (q *Queue[T]) Enqueue(data T) {
	for !q.stack1.IsEmpty() {
		q.stack2.Push(q.stack1.Peek())
		q.stack1.Pop()
	}

	q.stack1.Push(data)

	for !q.stack2.IsEmpty() {
		q.stack1.Push(q.stack2.Peek())
		q.stack2.Pop()
	}
}

func (q *Queue[T]) Dequeue() {
	q.stack1.Pop()
}

func (q *Queue[T]) Peek() T {
	return q.stack1.Peek()
}

func (q *Queue[T]) Size() int {
	return q.stack1.Size()
}

func (q *Queue[T]) IsEmpty() bool {
	return q.stack1.IsEmpty()
}

func (q *Queue[T]) Print() {
	q.stack1.Print()
}

func NewQueue[T comparable]() queue.IQueue[T] {
	return &Queue[T]{
		stack1: NewStack[T](),
		stack2: NewStack[T](),
	}
}
