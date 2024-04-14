package linkedlist

import (
	"fmt"
	. "github.com/abhisekp/go-repl/node"
	"strings"
)

type LinkedList[T comparable] struct {
	_    struct{}
	Head *Node[T]
	size int
}

type ILinkedList[T comparable] interface {
	Push(data T)
	Pop() T
	Size() int
	IsEmpty() bool
	Print()
}

type ILinkedListMethods[T comparable] interface {
	InsertAtHead(data T)
	InsertAtTail(data T)
	InsertAfterNode(node *Node[T])
	InsertBeforeNode(node *Node[T])
	InsertAtPosition(position int, data T)
	InsertAfterValue(value T, data T)
	InsertBeforeValue(value T, data T)

	Reverse()
	Search(data T) *Node[T]
	Traverse()
	TraverseReverse()
	IsLoopPresent() bool

	RemoveHead()
	RemoveTail()
	RemoveNode(node *Node[T])
	RemoveLoop()
	RemoveDuplicate()
	RemoveNthNodeFromEnd(n int)
	RemoveNthNodeFromStart(n int)
	RemoveNthNodeFromMiddle(n int)
	ReverseNthNodeFromStart(n int)
	ReverseNthNodeFromEnd(n int)
	ReverseNthNodeFromMiddle(n int)
	FindMiddleNode() *Node[T]
	FindLoop() *Node[T]
	FindLoopLength() int
	FindNthNodeFromEnd(n int) *Node[T]
	FindNthNodeFromStart(n int) *Node[T]
	FindNthNodeFromMiddle(n int) *Node[T]
}

func NewLinkedList[T comparable]() ILinkedList[T] {
	return &LinkedList[T]{}
}

func (ll *LinkedList[T]) Push(data T) {
	newNode := Node[T]{Data: data}

	defer func() {
		ll.size++
	}()

	if ll.Head == nil {
		ll.Head = &newNode
		return
	}

	curr := ll.Head
	for ; curr.Next != nil; curr = curr.Next {
	}

	curr.Next = &newNode
}

func (ll *LinkedList[T]) Pop() T {
	var empty T
	if ll.Size() == 0 {
		return empty
	}

	defer func() {
		ll.size--
	}()

	if ll.Size() == 1 {
		popped := ll.Head.Data
		ll.Head = nil
		return popped
	}

	curr := ll.Head
	for ; curr.Next.Next != nil; curr = curr.Next {
	}

	popped := curr.Next.Data
	curr.Next = nil

	return popped
}

func (ll *LinkedList[T]) Size() int {
	return ll.size
}

func (ll *LinkedList[T]) IsEmpty() bool {
	return ll.Size() == 0
}
func (ll *LinkedList[T]) Print() {
	var str strings.Builder

	str.WriteString("> ")
	for curr := ll.Head; curr != nil; curr = curr.Next {
		str.WriteString(fmt.Sprintf("%v -> ", curr.Data))
	}
	str.WriteString("|")

	fmt.Println(str.String())
}
