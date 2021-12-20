package main

import "fmt"

func message(text string, c chan<- string) {
	c <- text
}

func main() {
	c := make(chan string, 2)
	c <- "Message1"
	c <- "Message2"

	// capacity and length
	fmt.Println(len(c), cap(c))

	// close and range
	close(c)
	for message := range c {
		fmt.Println(message)
	}

	// select
	chan1 := make(chan string)
	chan2 := make(chan string)
	go message("message1", chan1)
	go message("message2", chan2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-chan1:
			fmt.Println("Email received from channel 1", m1)
		case m2 := <-chan2:
			fmt.Println("Email received from channel 2", m2)
		}
	}
}
