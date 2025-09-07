package main

import (
	"fmt"
	"project_1/vector"
)

func main(){

	fmt.Println("======================== 2D VECTOR ==============================")
	// Calculation of 2D vector is done
	v2d1 := vector.Vector2D{X:2, Y:4}
	v2d2 := vector.Vector2D{X:5, Y:6}

	v2d1.AddIn(v2d2)
	fmt.Println("Addition of 2D vector is : ",v2d1)
	v2d1.Subtract(v2d2)
	fmt.Println("Subtraction of 2D vector is : ",v2d1)
	dot := v2d1.Dot_Product(v2d2)
	fmt.Println("Dot product of 2D vector is : ",dot)

	fmt.Println("======================== 3D VECTOR ==============================")
	// Calculation of 3D vector is done
	v3d1 := vector.Vector3D{X:4, Y:6, Z:8}
	v3d2 := vector.Vector3D{X:5, Y:2, Z:5}

	v3d1.AddVector3D(v3d2)
	fmt.Println("Addition of 3D vector is : ",v3d1)
	v3d1.SubtractVector3D(v3d2)
	fmt.Println("Subtraction of 3D vector is : ",v3d1)
	dot = v3d1.DotVector3D(v3d2)
	fmt.Println("Dot product of 3D vector is : ",dot)
}