package main

import "fmt"

type Person struct {
	name string
	age  int
}

type BaseEmployee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	BaseEmployee
}
func (ftEmployee FullTimeEmployee) getMessage() string{
	return "Full time employee"
}

type TemporaryEmployee struct {
	Person
	BaseEmployee
	taxRate int
}
func (tEmployee TemporaryEmployee) getMessage() string{
	return "Temporary employee"
}

// INTERFACE AND IMPLICIT IMPLEMENT

type PrintInfo interface {
	getMessage() string
}
func GetMessage(p PrintInfo)  { // here param and interface merge
	fmt.Println(p.getMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Mar"
	ftEmployee.age = 28
	ftEmployee.id = 1
	fmt.Printf("%v\n", ftEmployee)
	tEmployee := TemporaryEmployee{}

	GetMessage(ftEmployee)
	GetMessage(tEmployee)

}
