package main

import (
	"fmt"
	"project_3/linkedlist"
)

func main(){
	fmt.Println("====================================")
	fmt.Println("Single linked list")
	linkedlist.Single_Linked_list()
	fmt.Println("")
	fmt.Println("====================================")
	fmt.Println("Doubly linked list")
	linkedlist.Doubly_Linkedlist()
	fmt.Println("")
	fmt.Println("=====================================")
	fmt.Println("Single circular linked list")
	linkedlist.Circular_linked_list()
	fmt.Println("")
	fmt.Println("=====================================")
	fmt.Println("Doubly circular linked list")
	linkedlist.Doubly_Circular_linkedlist()
}
