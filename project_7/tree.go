package main

import "fmt"

type Tree struct{
	Value 		int
	LeftNode 	*Tree
	RightNode 	*Tree
}

func main(){
	root := Tree{
				Value: 45,
				LeftNode: &Tree{Value: 34},
				RightNode: &Tree{Value: 56},
				}
	root.LeftNode.LeftNode = &Tree{Value: 24}
	root.LeftNode.RightNode = &Tree{Value: 30}

	root.RightNode.LeftNode = &Tree{Value: 50}
	root.RightNode.RightNode = &Tree{Value: 60}

	// Print Tree manually
	fmt.Println("Root Node : ",root.Value)
	fmt.Println("Left Node : ",root.LeftNode.Value)
	fmt.Println("Right Node : ",root.RightNode.Value)

	//Now leftNode  value
	fmt.Println("Left Node of left node : ",root.LeftNode.LeftNode.Value)
	fmt.Println("Right Node of left node : ",root.LeftNode.RightNode.Value)

	// Now RightNode value
	fmt.Println("left Node of Right node : ",root.RightNode.LeftNode.Value)
	fmt.Println("Right Node of Right node : ",root.LeftNode.LeftNode.Value)
}
