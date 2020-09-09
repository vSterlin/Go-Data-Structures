package main

import (
	"github.com/vSterlin/go_data_structures/linkedlist"
)

func main() {

	n := linkedlist.NewLinkedList()
	for i := 1; i < 10; i++ {
		n.AddToFront(i)

	}

	n.PrintList()

}
