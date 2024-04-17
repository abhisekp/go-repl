package doubleendedlinkedlist

import "testing"

func TestDoubleEndedLinkedList(t *testing.T) {
	t.Run("InsertAtHead", func(t *testing.T) {
		dll := NewDLL[int]()
		dll.InsertAtHead(1)
		dll.InsertAtHead(2)
		dll.InsertAtHead(3)

		if dll.Head().Data != 3 {
			t.Errorf("head() = %v, want %v", dll.Head().Data, 3)
		}
	})
}
