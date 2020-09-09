package linkedlist

import (
	"fmt"
	"strconv"
)

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func NewLinkedList() *LinkedList {
	l := new(LinkedList)
	l.Head = nil
	l.Tail = nil
	l.Size = 0
	return l
}

func (l *LinkedList) IsEmpty() bool {
	if l.Size == 0 {
		return true
	} else {
		return false
	}
}

func (l *LinkedList) AddToFront(value int) {

	nodeToAdd := new(Node)
	nodeToAdd.Data = value
	if l.IsEmpty() == true {
		l.Head = nodeToAdd
		l.Tail = nodeToAdd
	} else {
		nodeToAdd.Next = l.Head
		l.Head.Prev = nodeToAdd
		l.Head = nodeToAdd
	}
	l.Size++

}
func (l *LinkedList) AddToRear(value int) {
	nodeToAdd := new(Node)
	nodeToAdd.Data = value
	if l.IsEmpty() == true {
		l.Head = nodeToAdd
		l.Tail = nodeToAdd
	} else {
		nodeToAdd.Prev = l.Tail
		l.Tail.Next = nodeToAdd
		l.Tail = nodeToAdd
	}

	l.Size++
}

func (l *LinkedList) AddAtIndex(value int, index int) {
	if index <= 0 {
		l.AddToFront(value)
	} else if index >= l.Size {
		l.AddToRear(value)
	} else {
		current := l.Head
		nodeToAdd := new(Node)
		nodeToAdd.Data = value
		i := 0
		for i < index {
			current = current.Next
			i++
		}
		nodeToAdd.Prev = current.Prev
		nodeToAdd.Next = current
		current.Prev.Next = nodeToAdd
		current.Prev = nodeToAdd
		l.Size++
	}
}

func (l *LinkedList) GetFront() int {
	return l.Head.Data
}
func (l *LinkedList) GetRear() int {
	return l.Tail.Data
}

func (l *LinkedList) GetAtIndex(index int) int {
	i := 0
	current := l.Head
	for i < index {
		current = current.Next
		i++
	}
	return current.Data
}

func (l *LinkedList) DeleteFront() {
	if l.IsEmpty() == true {
		return
	} else {
		l.Head = l.Head.Next
		l.Head.Prev = nil
	}
	l.Size--
}

func (l *LinkedList) DeleteRear() {
	if l.IsEmpty() == true {
		return
	} else {
		l.Tail = l.Tail.Prev
		l.Tail.Next = nil
	}
	l.Size--
}

func (l *LinkedList) DeleteAtIndex(index int) {
	if index <= 0 {
		l.DeleteFront()
	} else if index >= l.Size {
		l.DeleteRear()
	} else {
		current := l.Head
		i := 0
		for i < index {
			current = current.Next
			i++
		}
		current.Prev.Next = current.Next
		current.Next.Prev = current.Prev
		l.Size--
	}
}

func (l *LinkedList) PrintList() {
	outputString := "{ "
	current := l.Head
	for current != nil {
		outputString += strconv.Itoa(current.Data) + ", "
		current = current.Next
	}
	outputString = outputString[0 : len(outputString)-2]
	outputString += " }"
	fmt.Println(outputString)
}
