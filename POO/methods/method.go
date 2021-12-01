package main

import "fmt"

type Employee struct {
	id   int
	name string
}

// RECEIVERS FUNCTIONS

func (e *Employee) setId(id int) {
	e.id = id
}
func (e *Employee) setName(name string) {
	e.name = name
}
func (e *Employee) getId() int {
	return e.id
}
func (e *Employee) getName() string {
	return e.name
}

func main() {
	employee := Employee{}
	fmt.Println("new employee", employee)
	employee.setId(23)
	employee.setName("wilmer")
	fmt.Printf("new employee %s with code %d\n", employee.getName(), employee.getId())
}
