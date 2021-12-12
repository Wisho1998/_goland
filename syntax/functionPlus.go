package main

import (
	. "fmt"
)

// variadic function
func printNames(names ...string) {
	for _, v := range names {
		Println("value : ", v)
	}
}

// name return parameter
func getMultiple(x int) (double int, triple int, quad int) {
	// return 2 * x, 3 * x, 4 * x
	double = 2 * x
	triple = 3 * x
	quad = 4 * x
	return
}

func main() {
	names := []string{"Henry", "Miguel1", "Miguel2", "Edgar"}
	printNames(names...)

	Println(getMultiple(5))
}
