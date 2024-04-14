package sortlistusingstack

import "github.com/abhisekp/go-repl/stackusingqueue"

func Run() {
	SortListUsingStack([]int{1, 9, 3, 6, 5, 2, 0})
}

func SortListUsingStack[T comparable](unsortedList []T) (sortedList []T) {
	s1 := stackusingqueue.NewStack[T]()
	s2 := stackusingqueue.NewStack[T]()

	for _, num := range unsortedList {
		if s1.IsEmpty() || s1.Peek() == num {
			s1.Push(num)
			continue
		}

		for s1.Peek() > num {
			s2.Push(s1.Peek())
			s1.Pop()
		}
	}

	return
}
