package vector


type Vector2D struct {
	X float32
	Y float32
}

func (v1 *Vector2D) AddIn(v2 Vector2D) {
	//  Here we Add two vector
	v1.X += v2.X // It means v1 = v1.X + v2.X
	v1.Y += v2.Y // It means v1 = v1.y + v2.y
}

func (v1 *Vector2D) Subtract(v2 Vector2D) {
	// Here we Subtract two vector
	v1.X -= v2.X
	v1.Y -= v2.Y
}


// Here we use value receiver instead of pointer receiver because the function 
// return a single vaue so we don't need to change the original vector which could be misleading

func (v1 Vector2D) Dot_Product(v2 Vector2D) float32 {
	// Here we Do Dot product
	return (v1.X * v2.X)+(v1.Y * v2.Y)
}


