package stack

import "testing"

func TestStack(t *testing.T) {
	t.Run("Push and Pop", func(t *testing.T) {
		s := NewStack[int]()
		s.Push(1)
		if s.Peek() != 1 {
			t.Errorf("Expected 1, got %v", s.Peek())
		}

		s.Push(2)
		if s.Peek() != 2 {
			t.Errorf("Expected 2, got %v", s.Peek())
		}

		s.Push(3)
		if s.Peek() != 3 {
			t.Errorf("Expected 3, got %v", s.Peek())
		}

		s.Pop()
		if s.Peek() != 2 {
			t.Errorf("Expected 2, got %v", s.Peek())
		}

		s.Pop()
		if s.Peek() != 1 {
			t.Errorf("Expected 1, got %v", s.Peek())
		}

		s.Pop()
		if !s.IsEmpty() {
			t.Errorf("Expected true, got %v", s.IsEmpty())
		}

		if s.Peek() != 0 {
			t.Errorf("Expected 0, got %v", s.Peek())
		}

		s.Pop()
		if s.Size() != 0 {
			t.Errorf("Expected 0, got %v", s.Size())
		}
	})
}
