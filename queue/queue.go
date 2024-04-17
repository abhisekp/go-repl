package queue

import (
	"errors"
	"fmt"
	"strings"
)

type IQueue[T any] interface {
	Enqueue(data T)
	Dequeue()
	Peek() T
	Size() int
	IsEmpty() bool

	Print()
}

type Queue[T any] struct {
	_      struct{}
	_queue []T
}

type QueueOption struct {
	_ struct{}
}

type QueueFn[T any] func(queue *Queue[T])

func WithCapacity[T any](capacity int) QueueFn[T] {
	if capacity < 0 {
		panic(errors.New("Capacity must be greater than 0"))
	}

	return func(queue *Queue[T]) {
	}
}

func NewQueue[T any](options ...QueueFn[T]) IQueue[T] {
	queue := &Queue[T]{}

	for _, optionFn := range options {
		optionFn(queue)
	}

	return queue
}

func (q *Queue[T]) Enqueue(data T) {
	q._queue = append(q._queue, data)
}

func (q *Queue[T]) Dequeue() {
	if q.IsEmpty() {
		return
	}

	q._queue = q._queue[1:]
}

func (q *Queue[T]) Peek() T {
	var empty T

	if q.IsEmpty() {
		return empty
	}

	return q._queue[0]
}

func (q *Queue[T]) Size() int {
	return len(q._queue)
}

func (q *Queue[T]) Print() {
	var str strings.Builder
	str.WriteString("> ")
	for curr := range q._queue {
		str.WriteString(fmt.Sprintf("%v -> ", curr))
	}
	str.WriteString("|")

	fmt.Println(str.String())
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q._queue) == 0
}
