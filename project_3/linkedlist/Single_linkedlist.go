package linkedlist

import "fmt"

type Node struct {
	data  int
	node  *Node
}

func Single_Linked_list(){

	// &Node{data: 10} := we use this because 
	/*					=>The & operator means “give me the address of this struct.”
						  Instead of holding the whole struct as a value, we now hold a pointer to it.
						  Why? Because linked lists work by connecting nodes using addresses (pointers), not copies.
	*/
	node1 := &Node{data: 10}
	node2 := &Node{data: 20}
	node3 := &Node{data: 30}
	node4 := &Node{data: 40}

	node1.node = node2
	node2.node = node3
	node3.node = node4

	current := node1
	for current != nil {
		fmt.Println("The value is : ", current.data)
		current = current.node
	}
}
