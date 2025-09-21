package main

import "fmt"

// Task struct with pointer references to dependencies
type Task struct{
	Name 	string
	Dependencies []*Task    		// Slice of pointer to dependent task
}

// AddDependency adds a dependency to a task
func (t *Task) AddDependency(dep *Task){
	t.Dependencies = append(t.Dependencies, dep)
}

// Display dependencies prints task dependencies
func (t *Task) Display(){
	fmt.Printf("Task %s depends on : ",t.Name)
	if len(t.Dependencies)== 0{
		fmt.Println("None")
		return
	}
	for _, dep := range t.Dependencies{
		fmt.Printf("%s", dep.Name)
	}
	fmt.Println()
}

func main(){
	// Create Task
	Build := &Task{Name: "Build"}
	Test := &Task{Name: "Test"}
	Deploy := &Task{Name: "Deploy"}

	// Define dependencies
	Test.AddDependency(Build)          // Test Depends on build
	Deploy.AddDependency(Test)		   // Deploy depends on Test

	Build.Display()
	Test.Display()
	Deploy.Display()

}