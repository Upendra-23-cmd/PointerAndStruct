package main

import "fmt"

func swap(a,b *int){
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}

func swap_value(a,b int){
	a, b = b , a
	c := &a
	d := &b
	fmt.Println(c , "  ",d)
}

func main(){

	a := 34
	b := 45
	fmt.Printf("the value of a is %d and b is %d\n",a,b)
	swap(&a,&b)
	fmt.Printf("the value of a is %d and b is %d\n",a,b)
	c := &a
	d := &b
	fmt.Println(c , "  ",d)
	swap_value(a,b)

}
