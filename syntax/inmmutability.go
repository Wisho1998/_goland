package main

import "fmt"

func main(){
	myArray := make([]bool, 3)
	//var myArray [3]bool // Array too

	// -. If only cut a slice the reference(pointers) persists
	//temp := myArray[:1]
	//temp[0] = true
	//fmt.Println(myArray) // [true false false]
	//fmt.Println(temp) // [true]


	// -. Instead, if we need an immutable copy we use "copy"
	temp := make([]bool, 1)
	copy(temp,myArray[:1])
	temp[0] = true

	fmt.Println(myArray) // [false false false]
	fmt.Println(temp) // [true]
}