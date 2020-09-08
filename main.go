package main

import (
	"fmt"

	"github.com/vSterlin/go_data_structures/linkedlist"
)

func main() {

	n := linkedlist.NewLinkedList()
	for i := 1; i < 10; i++ {
		n.AddToFront(i)

	}
	n.DeleteFront()
	n.DeleteFront()
	n.DeleteRear()
	n.DeleteRear()
	n.DeleteRear()

	fmt.Printf("Front: %d\nRear: %d\n", n.GetFront(), n.GetRear())

}
