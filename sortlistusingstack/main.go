package sortlistusingstack

import (
	"cmp"
	"github.com/abhisekp/go-repl/stack"
)

func Run() {
	SortListUsingStack([]int{1, 9, 3, 6, 5, 2, 0})
}

func SortListUsingStack[T cmp.Ordered](unsortedList []T) (sortedList []T) {
	s1 := stack.NewStack[T]()
	s2 := stack.NewStack[T]()

	for _, num := range unsortedList {
		if s1.IsEmpty() || s1.Peek() == num {
			s1.Push(num)
			continue
		}

		for num < s1.Peek() {
			s2.Push(s1.Peek())
			s1.Pop()
		}

		for !s2.IsEmpty() {
			s1.Push(s2.Peek())
			s2.Pop()
		}
	}

	sortedList = make([]T, 0, len(unsortedList))
	for !s1.IsEmpty() {
		sortedList = append(sortedList, s1.Peek())
		s1.Pop()
	}
	return
}
