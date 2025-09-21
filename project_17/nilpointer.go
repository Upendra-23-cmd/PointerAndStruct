package main

import "fmt"

// Person struct with a pointer field
type Person struct{
	Name		string
	Age 		*int 			// age is a pointer to int
}

// PrintDetails handles nil pointer safely
func (p *Person) PrintDetails(){
	fmt.Printf("Name :%s\n", p.Name)
	if p.Age == nil {
		fmt.Println("Age : (not provided)")     // safe handleing 
	} else {
		fmt.Printf("Age : %d\n", *p.Age)		// deferences safely
	}
}

func main() {
	// Case 1: person nil pointer (no age)
	p1 := &Person{ Name: "Alice",Age: nil}
	p1.PrintDetails()

	// Case 2: person with age value 
	ageVal := 25
	p2 := &Person{Name: "Bob", Age : &ageVal}
	p2.PrintDetails()

	// Case 3: Nil struct pointer itself
	var p3 *Person = nil
	if p3 == nil {
	fmt.Println("p3 is nil, cannot call methods safely ")
	 } else {
		p3.PrintDetails()
	 }
}