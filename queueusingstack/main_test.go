package queueusingstack

import "testing"

func TestQueueUsingStack(t *testing.T) {
	t.Run("Enqueue and Dequeue", func(t *testing.T) {
		q := NewQueue[int]()

		q.Enqueue(1)
		if q.Peek() != 1 {
			t.Errorf("Expected 1, got %v", q.Peek())
		}

		q.Enqueue(2)
		if q.Peek() != 1 {
			t.Errorf("Expected 1, got %v", q.Peek())
		}

		q.Enqueue(3)
		if q.Peek() != 1 {
			t.Errorf("Expected 1, got %v", q.Peek())
		}

		q.Dequeue()
		if q.Peek() != 2 {
			t.Errorf("Expected 2, got %v", q.Peek())
		}

		if q.Size() != 2 {
			t.Errorf("Expected 2, got %v", q.Size())
		}

		q.Dequeue()
		if q.Peek() != 3 {
			t.Errorf("Expected 3, got %v", q.Peek())
		}

		q.Dequeue()
		if !q.IsEmpty() {
			t.Errorf("Queue should be empty")
		}

		q.Dequeue()
		if !q.IsEmpty() {
			t.Errorf("Queue should be empty")
		}

		if q.Peek() != 0 {
			t.Errorf("Expected 0, got %v", q.Peek())
		}
	})

}
