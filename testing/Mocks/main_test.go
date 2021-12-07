package main

import "testing"

func TestGetFullTimeEmployeeById(t *testing.T) {
	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			1,
			"1",
			func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{Id: 1, Position: "CEO"}, nil
				}
				GetPersonByDNI = func(dni string) (Person, error) {
					return Person{Name: "Wilmer", Age: 23, DNI: }, nil
				}
			},
			FullTimeEmployee{},
		},
	}
}
