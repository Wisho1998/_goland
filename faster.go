package main

import (
	"fmt"
	"math/rand"
)

func doSomething(x int) int {
	return x
}

func manageTraffic(entrada <-chan int, salida chan<- int) {
	for c := range entrada {
		salida <- doSomething(c)
	}
	close(salida)
}

func main() {
	numbers := make(chan int)
	doubles := make(chan int)
	//traffic := make(chan struct{})

	go publicar(numbers)
	go manageTraffic(numbers, doubles)
	suscribir(doubles)
}

func suscribir(conTotal <-chan int) {
	for c := range conTotal {
		fmt.Printf("%#v ", c)
	}
}

func publicar(output chan<- int) {
	for i := 0; i < 1000; i++ {
		output <- getRandom()
	}
	close(output)
}

func getRandom() int {
	min := 1
	max := 3
	return rand.Intn(max-min) + min
}
