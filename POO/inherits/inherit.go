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
	extRate string
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Mar"
	ftEmployee.age = 28
	ftEmployee.id = 1
	fmt.Printf("%v\n", ftEmployee)
}
