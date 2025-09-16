package main

import "fmt"

// It represent one element in stack
type Node struct{
	Data int
	Value *Node							
}

// It represent the stack 
type Stack struct{
	Top *Node							// points to the top node
}

// Push add value to the stack
func (s *Stack) Push(value int){
	newNode := &Node{Data: value}
	newNode.Value = s.Top				// point to the old top node
	s.Top = newNode						// Update top to new node
}

// Pop removes the element from the top and point to old top value
func (s *Stack) Pop()(int,bool){
	if s.Top == nil{
		return 0,false
	}
	value := s.Top.Data
	s.Top = s.Top.Value
	return value,true
}

// Peek returns top value without removeing it
func (s *Stack) Peek()(int, bool){
	if s.Top == nil{
		return 0,false
	}
	return s.Top.Data,true
}

// IsEmpty check the stack is empty or not
func (s *Stack) IsEmpty()bool{
	return s.Top == nil 	
}

 func main(){
	Stack:= Stack{}

	Stack.Push(10)
	Stack.Push(20)
	Stack.Push(30)
	Stack.Push(40)

	val, _ := Stack.Peek()
	fmt.Println("Peek : ", val)

	for !Stack.IsEmpty(){
		val, _ := Stack.Pop()
		fmt.Println("Popped value : ", val)
	}
 }