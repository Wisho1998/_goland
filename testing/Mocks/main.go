package main

import "time"

type Person struct {
	DNI string
	Name string
	Age int
}

type Employee struct {
	Id int
	Position string
}

type FullTimeEmployee struct {
	Employee
	Person
}

var GetPersonByDNI = func(dni string)(Person, error){
	time.Sleep(time.Second * 5)
	// SELECT * FROM Persona WHERE
	return Person{}, nil
}

var GetEmployeeById = func(id int)(Employee, error){
	time.Sleep(time.Second * 5)
	// SELECT * FROM Persona WHERE
	return Employee{}, nil
}

func GetFullTimeEmployeeById(id int, dni string) (FullTimeEmployee, error){
	var ftEmployee FullTimeEmployee
	e,_ := GetEmployeeById(id)
	ftEmployee.Employee = e

	p,_ := GetPersonByDNI(dni)
	ftEmployee.Person = p

	return ftEmployee, nil
}