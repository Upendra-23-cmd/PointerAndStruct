package linkedlist

import "fmt"

type DNode struct{
	data int
	prev_node *DNode
	next_node *DNode
}

func Doubly_Linkedlist(){

	node1 := &DNode{data: 10}
	node2 := &DNode{data: 20}
	node3 := &DNode{data: 30}
	node4 := &DNode{data: 40}


	// connecting the linked list forward
	node1.next_node = node2
	node2.next_node = node3
	node3.next_node = node4

	// connecting backward
	node4.prev_node = node3
	node3.prev_node = node2
	node2.prev_node = node1

	fmt.Println("Forward traversal of linked list ")
	current := node1
	for current != nil {
		fmt.Println("The node are : ", current.data)
		current = current.next_node
	}

	fmt.Println("Backward traversal of linked list ")
	current = node4
	for current != nil {
		fmt.Println("The node are ", current.data)
		current = current.prev_node
	}

}