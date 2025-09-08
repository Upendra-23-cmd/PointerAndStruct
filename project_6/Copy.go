package main

import "fmt"


type Original struct{
	Name string
	Books []string
}

func main() {
	original := Original{
		Name: "Library",
		Books: []string{"go books","docker guide"},
	}

	shallow := original

	deep := Original{
		Name: shallow.Name,
		Books : append([]string(nil),shallow.Books...),
	}
	fmt.Println("")
	fmt.Println("The original struct is \n",original)
	fmt.Println("The shallow copy of  struct is \n",shallow)
	fmt.Println("The deep copy of  struct is \n",deep)

	fmt.Println("")
	fmt.Println("==============================================================")
	fmt.Println("")
	shallow.Books[0] = "Rust book"
	deep.Books[1]= "Go book"

	fmt.Println("The original struct is \n",original)
	fmt.Println("The updated shallow copy of  struct is \n",shallow)
	fmt.Println("The updated deep copy of struct is \n", deep)

	fmt.Println("")
}
