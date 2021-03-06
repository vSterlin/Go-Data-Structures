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
	n.DeleteRear()

	n.AddAtIndex(123, 2)
	n.DeleteAtIndex(2)

	fmt.Println(n.Size)
	n.PrintList()

}
