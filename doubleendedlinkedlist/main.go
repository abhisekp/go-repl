package doubleendedlinkedlist

import . "github.com/abhisekp/go-repl/node"

type DoubleEndedLinkedList[T comparable] struct {
	_    struct{}
	head *DLLNode[T]
	tail *DLLNode[T]
	size int
}

func (dll *DoubleEndedLinkedList[T]) InsertAtHead(data T) {
	newNode := &DLLNode[T]{Data: data}

	defer func() {
		dll.size++
	}()

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
		return
	}

	newNode.Next = dll.head
	dll.head.Prev = newNode
	dll.head = newNode
}

func (dll *DoubleEndedLinkedList[T]) InsertAtTail(data T) {
	newNode := &DLLNode[T]{Data: data}

	defer func() {
		dll.size++
	}()

	if dll.tail == nil {
		dll.tail = newNode
		dll.head = newNode
		return
	}

	dll.tail.Next = newNode
	newNode.Prev = dll.tail
	dll.tail = newNode
}

func (dll *DoubleEndedLinkedList[T]) Head() *DLLNode[T] {
	return dll.head
}

type IDoubleEndedLinkedList[T comparable] interface {
	InsertAtHead(data T)
	InsertAtTail(data T)
	Head() *DLLNode[T]
}

func NewDLL[T comparable]() IDoubleEndedLinkedList[T] {
	return &DoubleEndedLinkedList[T]{}
}
