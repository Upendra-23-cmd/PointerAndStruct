package main

import "fmt"


type Shape struct{
	Name 		string
	Edges		string
}

type Circle struct{
	Shape
	Radius    int
}

type Rectangle struct{
	Shape
	Length     int
	Breadth	   int
}

func main() {

	c := Circle{
		Shape : Shape{Name: "Circle"},
		Radius: 34,
	}

	r := Rectangle{
		Shape: Shape{Name: "Rectangle"},
		Length:  24,
		Breadth: 45,
	}

	fmt.Println("The information of the circle is : ",c)
	fmt.Println("The information of the rectangle is : ",r)
}