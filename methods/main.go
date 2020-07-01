package main

import (
	"fmt"
)

type Employee struct {
	Name string
}

// Method
// func (e *Employee) UpdateName(newName string) {
// 	e.Name = newName
// }

// Function
func UpdateName(employee *Employee, newName string) {
	employee.Name = newName
}

func (e *Employee) PrintName() {
	fmt.Println(e.Name)
}

func main() {
	var employee Employee
	employee.Name = "Elliot"
	// Method call
	// employee.UpdateName("Forbsey")
	// Function call
	UpdateName(employee, "Forbsey")
	//employee.PrintName()
}
