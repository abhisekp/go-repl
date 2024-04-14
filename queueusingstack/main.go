package queueusingstack

import (
	"github.com/abhisekp/go-repl/queue"
	. "github.com/abhisekp/go-repl/stack"
)

type Queue[T comparable] struct {
	_       struct{}
	_stack1 IStack[T]
	_stack2 IStack[T]
}

func (q *Queue[T]) Enqueue(data T) {
	for !q._stack1.IsEmpty() {
		q._stack2.Push(q._stack1.Peek())
		q._stack1.Pop()
	}

	q._stack1.Push(data)

	for !q._stack2.IsEmpty() {
		q._stack1.Push(q._stack2.Peek())
		q._stack2.Pop()
	}
}

func (q *Queue[T]) Dequeue() {
	q._stack1.Pop()
}

func (q *Queue[T]) Peek() T {
	return q._stack1.Peek()
}

func (q *Queue[T]) Size() int {
	return q._stack1.Size()
}

func (q *Queue[T]) IsEmpty() bool {
	return q._stack1.IsEmpty()
}

func (q *Queue[T]) Print() {
	q._stack1.Print()
}

func NewQueue[T comparable]() queue.IQueue[T] {
	return &Queue[T]{
		_stack1: NewStack[T](),
		_stack2: NewStack[T](),
	}
}
