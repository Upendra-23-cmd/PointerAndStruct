package main

import "fmt"

// Here we create a structure of the student and information it need
type student struct{
	Name     string
	Roll_no  int
	grades   float32
}

// Here we created a structure of the class of the student in which they are
type class struct{
	Name  string
	studentA student
	studentB student
	studentC student
}

// If we want to change the grade of the student in a class
func (s *student)Change_grade(newGrade float32){
  s.grades = newGrade
}

// We display the details of the student 
func (c class)display(){
	fmt.Printf("The Student of class %s are \n",c.Name)
	fmt.Printf("The first student Name is %s his roll no is %d and his grades is %.2f\n",c.studentA.Name,c.studentA.Roll_no,c.studentA.grades)
	fmt.Printf("The second student Name is %s his roll no is %d and his grades is %.2f\n",c.studentB.Name,c.studentB.Roll_no,c.studentB.grades)
	fmt.Printf("The Third student Name is %s his roll no is %d and his grades is %.2f\n",c.studentC.Name,c.studentC.Roll_no,c.studentC.grades)
}
func main(){

	// Insert data of the students and in which class they are
	ClassA := class{
		Name: "Class A",
		studentA: student{Name: "Upendra", Roll_no: 73,grades: 93},
		studentB: student{Name: "Golu",Roll_no: 23,grades: 45},
		studentC: student{Name: "Arpit",Roll_no: 34,grades: 66},
	}
	ClassA.display()

	// change the grade
	ClassA.studentA.Change_grade(97)
	ClassA.studentB.Change_grade(67)
	ClassA.studentC.Change_grade(78)

	// display the detail of the class
	ClassA.display()
}