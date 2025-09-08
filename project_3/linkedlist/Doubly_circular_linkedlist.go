package linkedlist

import "fmt"

type DCnode struct{
	data       int
	prev_node  *DCnode
	next_node  *DCnode
}

func Doubly_Circular_linkedlist(){

	node1 := &DCnode{data: 10}
	node2 := &DCnode{data: 20}
	node3 := &DCnode{data: 30}
	node4 := &DCnode{data: 40}

	// connecting forward node
	node1.next_node = node2
	node2.next_node = node3
	node3.next_node = node4
	node4.next_node = node1

	// forward travesal of linked list
	fmt.Println("Forward traversal of doubly circular linked list")
	current := node1
	count := 0
	for count <10{
		fmt.Println("The data is : ", current.data)
		current = current.next_node
		count++
	}

	// connecting backward node
	node4.prev_node = node3
	node3.prev_node = node2
	node2.prev_node = node1
	node1.prev_node = node4

	// Backward traversal of linked list
	fmt.Println("Backward traversal of doubly circular linked list")
	count = 10
	for count >0{
		fmt.Println("The data is : ", current.data)
		current = current.prev_node
		count--
	}
}
