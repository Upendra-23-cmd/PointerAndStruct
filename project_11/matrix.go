package main

import "fmt"

type Matrix struct{
	Row int
	Col int
	Data *[]float64
}

func NewMatrix(row, col int)*Matrix{
	data := make([]float64, row*col)
	return &Matrix{Row: row, Col: col, Data: &data}
}


func (m *Matrix) Set(row, col int, value float64){
	index := row*m.Col + col
	(*m.Data)[index] = value
}


// Get value
func (m *Matrix) Get(row, col int) float64 {
	index := row*m.Col + col
	return (*m.Data)[index]
}

func main() {
	// Create 2x3 matrix
	mat := NewMatrix(2, 3)

	// Set values
	mat.Set(0, 0, 1)
	mat.Set(0, 1, 2)
	mat.Set(0, 2, 3)
	mat.Set(1, 0, 4)
	mat.Set(1, 1, 5)
	mat.Set(1, 2, 6)

	// Print
	fmt.Println("Matrix:")
	for r := 0; r < mat.Row; r++ {
		for c := 0; c < mat.Col; c++ {
			fmt.Print(mat.Get(r, c), " ")
		}
		fmt.Println()
	}
}