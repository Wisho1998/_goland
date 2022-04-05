package main

import (
	"fmt"
	"time"
)

func main() {
	sema := make(chan struct{}, 3)
	canal := make(chan int)

	for i := 0; i < 10000; i++ {
		go func(i int) {
			sema <- struct{}{} // adquiere un token
			time.Sleep(1 * time.Second)
			canal <- i
			<-sema // devuelve el token
		}(i)
	}

	for i := 0; i < 10000; i++ {
		fmt.Printf("%v ",<-canal)
	}
}
