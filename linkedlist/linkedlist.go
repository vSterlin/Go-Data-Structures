package linkedlist

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

func (l *LinkedList) GetFront() int {
	return l.Head.Data
}
func (l *LinkedList) GetRear() int {
	return l.Tail.Data
}
