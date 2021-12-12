package main

import (
	"fmt"
	myPackage "github.com/Wisho1998/accessModifiers"
	"math/rand"
	"time"
)

func getRandom() int {
	min := 1
	max := 100
	return rand.Intn(max-min) + min
}

func main() {
	//	TOUR OF GO
	//https://go.dev/tour/list

	//PRINT MESSAGES IN CONSOLE
	fmt.Println()
	m := fmt.Sprintf("")
	fmt.Printf("%s", m)

	// DECLARE VARIABLES
	var zeroValue float64 // zero value
	var age int
	age = 23
	age2 := 23
	fmt.Println("VARIABLES", zeroValue, age, age2)

	// RANGE TYPES DATA
	// https://codingornot.com/02-go-to-go-sintaxis-tipo-de-datos-y-palabras-reservadas

	//ARRAYS
	// fixed size
	var myArray [4]bool
	fmt.Println("ARRAY ", myArray, len(myArray), cap(myArray))

	//SLICES
	// dynamic size
	mySlice := []int{1, 2, 3}
	mySlice = append(mySlice, 4)
	mySlice = append(mySlice, 5)
	fmt.Println("SLICE and slicing slice[min:max]", mySlice[2:len(mySlice)-1])

	//MAPS
	// associative, efficient search, untidy iterative
	myMap := make(map[string]string)
	myMap["name"] = "Wilmer"
	fmt.Println("MAP", myMap)

	//LOOPS
	// for conditional with range statement
	for i, v := range mySlice {
		fmt.Println("LOOP FOR RANGE ", i, " : ", v)
	}
	// for - while
	counter := 0
	for counter < 10 {
		fmt.Printf("LOOP FOR WHILE %d ,", counter)
		counter++
	}
	// for - forever
	//for {
	//	fmt.Println("Kill me")
	//}

	//CONDITIONALS
	//IF
	number1 := getRandom()
	modulo1 := number1 % 2
	if modulo1 == 0 {
		fmt.Printf("\nIF | %d is Even\n", number1)
	} else {
		fmt.Printf("IF | %d is Odd\n", number1)

	}
	//SWITCH1
	number2 := getRandom()
	switch modulo2 := number2 % 2; modulo2 {
	case 0:
		fmt.Printf("SWITCH1 | %d is Even\n", number2)
	default:
		fmt.Printf("SWITCH1 | %d is Odd\n", number2)
	}
	//SWITCH2
	number3 := getRandom()
	switch {
	case number3 == 50:
		fmt.Printf("SWITCH2 | %d is Equal to %d\n", number3, 50)
	case number3 < 50:
		fmt.Printf("SWITCH2 | %d is Smaller than %d\n", number3, 50)
	default:
		fmt.Printf("SWITCH2 | %d is Greater than %d\n", number3, 50)
	}

	// STRUCTS
	myStruct := myPackage.MyStructPublic{}
	fmt.Println("STRUCTS", myStruct)

	// POINTERS https://ed.team/blog/que-son-los-punteros-en-go
	// used to store the memory address
	v := 19
	var p1 *int
	var p2 = new(int)
	p3 := &v
	fmt.Printf("p1: %T , p2 : %T and p3: %T\n", p1, p2, p3)
	fmt.Printf("p3: %d \n", *p3) // dereference operator

	// CHANNEL AND CONCURRENCY
	c := make(chan string)
	go func() { // anonymous function
		fmt.Println("Starting function")
		time.Sleep(5 * time.Second)
		c <- "end channel"
	}()
	<-c
}
