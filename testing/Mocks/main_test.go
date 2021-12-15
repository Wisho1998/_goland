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
					return Employee{Id: 1, Position: "dev"}, nil
				}
				GetPersonByDNI = func(dni string) (Person, error) {
					return Person{Name: "Wilmer", Age: 23, DNI: "1"}, nil
				}
			},
			FullTimeEmployee{
				Person: Person{
					Age:  23,
					DNI:  "1",
					Name: "Wilmer",
				},
				Employee: Employee{
					Id:       1,
					Position: "dev",
				},
			},
		},
	}
	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByDNI := GetPersonByDNI
	for _, item := range table {
		item.mockFunc()
		ft, err := GetFullTimeEmployeeById(item.id, item.dni)
		if err != nil {
			t.Errorf("Error when getting Employee")
		}
	}
}
