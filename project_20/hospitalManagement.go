package main

import "fmt"

// Write a program to Build a hospital patient management system using structs and pointers.

type Patient struct{
	Id   int
	Name string
	Age  int
	Disease string
	Next *Patient
}

type Hospital struct{
	Name  string
 	Patient *Patient
}

func (h *Hospital) NewPatient(disease,name string,id,age int){
	newPatient := &Patient{
		Id: id,
		Name: name,
		Age: age,
		Disease: disease,
	}

	h.Patient = newPatient
	fmt.Println("New patient registered \n details are : ",newPatient)

}

func (h *Hospital) findPatient(id int) *Patient{
	current := h.Patient
	for current != nil {
		if current.Id == id {
			return current
		}
		current = current.Next
	}
	return nil
}

func (h *Hospital) Updatedetail(id int , disease string){
	Patient := h.findPatient(id)
	if Patient != nil {
		Patient.Disease = disease
		fmt.Printf("Patient %s details are update to %s\n",Patient.Name,disease)
	}else{
		fmt.Println("Patient not found")
	}
}

func (h *Hospital) DisplayPatient(){
	if h.Patient == nil{
		fmt.Println("NO patient in the hospital")
	}
	current := h.Patient
	for current != nil {
		fmt.Printf("ID : %d | Name : %s | Age : %d | Disease : %s \n",current.Id,current.Name,current.Age,current.Disease)
		current = current.Next
	}
}

func main(){
	Hospital := Hospital{Name: "City hospital"}

	Hospital.NewPatient("fever" , "upendra" ,1,22)
	Hospital.NewPatient("cough","jhon", 2, 34)

	Hospital.DisplayPatient()

	Hospital.findPatient(1)
	
	
	Hospital.Updatedetail(2,"asthma")
}