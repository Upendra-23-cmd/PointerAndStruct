package main

import "fmt"

// Represent each node in a queue
type Node struct{
	Data int
	Value *Node
}

// Queue represent the pointer based queue
type Queue struct{
	Front *Node
	Rear  *Node
}

func (s *Queue) EnQueue(value int){
	newNode := &Node{Data: value,Value: nil}

	// If queue is empty
	if s.Rear == nil{
		s.Front = newNode
		s.Rear  = newNode
	}

	// Add to the rear and update rear pointer
	s.Rear.Value = newNode
	s.Rear = newNode
}

func (s *Queue) DeQueue()(int, bool){
	if s.Front == nil {
		return 0, false
	}

	// Store front data and move front pointer
	value := s.Front.Data
	s.Front = s.Front.Value

	// If Queue becomes empty , set rear nil
	if s.Front == nil {
		s.Rear = nil
	}
	return value, true
}

// Peek at the top element in the queue without removing it
func (s *Queue) Peek()(int,bool){
	if s.Front == nil {
		return 0,false
	}
	value := s.Front.Data
	return value, true
}

// check the queue is empty or not
func (s *Queue) IsEmpty() bool {
	return s.Front == nil
}

func (s *Queue) Display(){
	if s.Front == nil {
		fmt.Println("The Queue is empty")
	}

	fmt.Println("The queue elements are : ")
	for temp := s.Front; temp != nil ; temp = temp.Value{
		fmt.Printf("%d\n",temp.Data)
	}
	fmt.Println()
}

func main() {
	q := &Queue{}

	q.EnQueue(10)
	q.EnQueue(20)
	q.EnQueue(30)
	q.EnQueue(40)
	q.EnQueue(50)
	q.EnQueue(60)

	q.Display()

	val, ok := q.DeQueue()
	if ok {
		fmt.Println("Dequeued value : ",val)
	}

	peekValue, ok := q.Peek()
	if ok {
		fmt.Println("Peeked value is : ", peekValue)
	}

	q.Display()

}