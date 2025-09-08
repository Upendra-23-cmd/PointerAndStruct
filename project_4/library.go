package main

import "fmt"

type Book struct{
	Name string
	Author string
	UID int
}

func Book_Library(){

	book1 := &Book{
			Name: "The Magica",
			Author: "Upendra",
			UID: 1,
			}
	
	book2 := &Book{
			Name: "The lost",
			Author: "golu",
			UID: 2,
			}

	book3 := &Book{
			Name: "The found",
			Author: "Upendra",
			UID: 3,
			}
	// with using pointer
	Lib :=[]*Book{book1,book2,book3}
	for _, b := range Lib {
		fmt.Printf("Name : %s\n  Author : %s\n UID : %d\n",b.Name,b.Author,b.UID)
	}

	// without using the pointer 
	Lib2 := []Book{{"The magica ","Upendra",1},{"The lost","golu",2}}
	for _, b:= range Lib2 {
		fmt.Printf("Name : %s\n  Author : %s\n UID : %d\n",b.Name,b.Author,b.UID)
	}	

	// Updating the value of the struct
	book1.UID = 20

	for _, b := range Lib {
		fmt.Printf("Name : %s\n  Author : %s\n UID : %d\n",b.Name,b.Author,b.UID)
	}

	// without using the pointer 
	for _, b:= range Lib2 {
		fmt.Printf("Name : %s\n  Author : %s\n UID : %d\n",b.Name,b.Author,b.UID)
	}	

}

func main(){
	Book_Library()
}