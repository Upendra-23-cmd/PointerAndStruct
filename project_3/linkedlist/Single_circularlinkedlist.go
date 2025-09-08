package linkedlist

import "fmt"

type SCnode struct{
	data int
	next_node *SCnode
}

func Circular_linked_list(){

	node1 := &SCnode{data: 10}
	node2 := &SCnode{data: 20}
	node3 := &SCnode{data: 30}
	node4 := &SCnode{data: 40}

	// here we connect all the node to the next node
	node1.next_node = node2
	node2.next_node = node3
	node3.next_node = node4
	node4.next_node = node1

	current := node1
	count := 0
	for count < 10 {
		fmt.Println("The data is ",current.data)
		current = current.next_node
		count++
	}
}