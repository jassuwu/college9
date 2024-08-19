package six

import (
	"fmt"
)

type Node struct {
	Value int
	Prev  *Node
	Next  *Node
}

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

func (dll *DoublyLinkedList) InsertAtEnd(value int) {
	newNode := &Node{Value: value}
	if dll.Head == nil {
		dll.Head = newNode
		dll.Tail = newNode
	} else {
		dll.Tail.Next = newNode
		newNode.Prev = dll.Tail
		dll.Tail = newNode
	}
}

func (dll *DoublyLinkedList) InsertAtBeginning(value int) {
	newNode := &Node{Value: value}
	if dll.Head == nil {
		dll.Head = newNode
		dll.Tail = newNode
	} else {
		newNode.Next = dll.Head
		dll.Head.Prev = newNode
		dll.Head = newNode
	}
}

func (dll *DoublyLinkedList) DeleteByValue(value int) {
	current := dll.Head
	for current != nil {
		if current.Value == value {
			if current.Prev != nil {
				current.Prev.Next = current.Next
			} else {
				dll.Head = current.Next
			}
			if current.Next != nil {
				current.Next.Prev = current.Prev
			} else {
				dll.Tail = current.Prev
			}
			return
		}
		current = current.Next
	}
}

func (dll *DoublyLinkedList) PrintList() string {
	var result string
	current := dll.Head
	for current != nil {
		result += fmt.Sprintf("%d ", current.Value)
		current = current.Next
	}
	return result
}
