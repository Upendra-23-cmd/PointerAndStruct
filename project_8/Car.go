package main

import "fmt"

type Engine struct {
	Name      string
	Power     int
	Engine_No int
}

type Car struct {
	Name   string
	Engine *Engine
}

func main() {
	engine := &Engine{Name: "V58", Power: 2500, Engine_No: 345}
	car := &Car{Name: "BMW", Engine: engine}

	fmt.Printf("The information of car %s is \n", car.Name)
	fmt.Printf("The engine name of the car is %s ,its power is %d HP and the engine no. is %d\n", car.Engine.Name, car.Engine.Power, car.Engine.Engine_No)

	car.Engine.Engine_No = 2345
	fmt.Println("The updated value of the car engine is :- ",car.Engine.Engine_No)
}
