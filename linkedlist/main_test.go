package linkedlist

import "testing"

func TestLinkedList(t *testing.T) {
	t.Run("Push and Pop", func(t *testing.T) {
		ll := NewLinkedList[int]()
		ll.Push(1)
		if ll.Pop() != 1 {
			t.Errorf("Expected 1, got %v", ll.Pop())
		}

		ll.Push(2)
		if ll.Pop() != 2 {
			t.Errorf("Expected 2, got %v", ll.Pop())
		}

		ll.Push(3)
		if ll.Pop() != 3 {
			t.Errorf("Expected 3, got %v", ll.Pop())
		}

		if !ll.IsEmpty() {
			t.Errorf("Expected true, got %v", ll.IsEmpty())
		}

		ll.Push(4)
		ll.Push(5)
		ll.Push(6)

		if ll.Pop() != 6 {
			t.Errorf("Expected 6, got %v", ll.Pop())
		}
	})

	t.Run("Check loop exists", func(t *testing.T) {
		testCases := []struct {
			input    []int
			expected bool
		}{
			{
				input:    []int{1, 2, 3, 4, 5},
				expected: false,
			},
			{
				input:    []int{1, 2, 3, 4, 5, 1},
				expected: true,
			},
			{
				input:    []int{1, 2, 3, 4, 5, 4},
				expected: true,
			},
		}

		for _, tc := range testCases {
			ll := NewLinkedList[int](tc.input)
			ll.Print()
			if ll.IsLoopPresent() != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, ll.IsLoopPresent())
			}
		}
	})
}
