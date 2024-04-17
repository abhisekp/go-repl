package linkedlist

import (
	"cmp"
	"fmt"
	. "github.com/abhisekp/go-repl/node"
	"strings"
)

type LinkedList[T cmp.Ordered] struct {
	_    struct{}
	head *Node[T]
	size int
}

type ILinkedList[T cmp.Ordered] interface {
	Head() *Node[T]
	Push(data T)
	Pop() T
	Size() int
	IsEmpty() bool
	Print()
	IsLoopPresent() bool

	InsertAtHead(data T)
	RemoveAtHead()
	InsertAtTail(data T)

	Search(data T) *Node[T]

	Traverse(func(data T, nodes ...*Node[T]) bool)

	Tail() *Node[T]
}

type ILinkedListMethods[T cmp.Ordered] interface {
	InsertAfterNode(node *Node[T])
	InsertBeforeNode(node *Node[T])
	InsertAtPosition(position int, data T)
	InsertAfterValue(value T, data T)
	InsertBeforeValue(value T, data T)

	Reverse()
	Traverse()
	TraverseReverse()

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

func NewLinkedList[T cmp.Ordered](list ...[]T) ILinkedList[T] {
	newLinkedList := LinkedList[T]{}

	var fromList []T
	if len(list) > 0 {
		fromList = list[0]
		lastNode := newLinkedList.Tail()

		for _, data := range fromList {
			foundNode := newLinkedList.Search(data)
			// Loop Create
			if foundNode != nil && lastNode != nil {
				lastNode.Next = foundNode
				break
			}
			newLinkedList.Push(data)
			lastNode = newLinkedList.Tail()
		}
	}

	return &newLinkedList
}

func (ll *LinkedList[T]) Push(data T) {
	newNode := Node[T]{Data: data}

	defer func() {
		ll.size++
	}()

	if ll.head == nil {
		ll.head = &newNode
		return
	}

	curr := ll.head
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
		popped := ll.head.Data
		ll.head = nil
		return popped
	}

	curr := ll.head
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
	ll.Traverse(func(data T, nodes ...*Node[T]) bool {
		curr := nodes[0]
		str.WriteString(fmt.Sprintf("%v -> ", curr.Data))
		return true
	})
	str.WriteString("|")

	fmt.Println(str.String())
}

func (ll *LinkedList[T]) InsertAtHead(data T) {
	newNode := Node[T]{Data: data, Next: ll.head}

	ll.head = &newNode
	ll.size++
}

func (ll *LinkedList[T]) RemoveAtHead() {
	if ll.head == nil {
		return
	}

	ll.head = ll.head.Next
	ll.size--
}

func (ll *LinkedList[T]) Head() *Node[T] {
	return ll.head
}

func (ll *LinkedList[T]) InsertAtTail(data T) {
	ll.Push(data)
}

func (ll *LinkedList[T]) IsLoopPresent() bool {
	for slow, fast := ll.head, ll.head; ; {
		if fast == nil || fast.Next == nil || fast.Next.Next == nil {
			return false
		}

		slow, fast = slow.Next, fast.Next.Next

		if fast == slow {
			return true
		}
	}

	return true
}

func (ll *LinkedList[T]) Search(data T) *Node[T] {
	foundNode := (*Node[T])(nil)

	ll.Traverse(func(d T, node ...*Node[T]) bool {
		newNode := node[0]
		if data == d {
			foundNode = newNode
			return false
		}

		return true
	})

	return foundNode
}

func (ll *LinkedList[T]) Traverse(iter func(data T, nodes ...*Node[T]) bool) {

	curr := ll.head

	for range ll.Size() {
		if curr == nil || !iter(curr.Data, curr) {
			break
		}
		curr = curr.Next
	}

}

func (ll *LinkedList[T]) Tail() *Node[T] {
	curr := ll.head

	for ; curr != nil && curr.Next != nil; curr = curr.Next {
	}

	return curr
}
