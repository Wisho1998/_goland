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
	for _, test := range table {
		test.mockFunc()
		ft, err := GetFullTimeEmployeeById(test.id, test.dni)
		if err != nil {
			t.Errorf("Error when getting Employee")
		}

		if ft.Age != test.expectedEmployee.Age {
			t.Errorf("Error, got %d expected %d", ft.Age, test.expectedEmployee.Age)
		}
	}

	GetEmployeeById = originalGetEmployeeById
	GetPersonByDNI = originalGetPersonByDNI
}
