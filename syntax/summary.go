package main

import (
	"fmt"
	"time"
)

func main() {
	var age int
	age = 23

	age2 := 23
	fmt.Println(age, age2)

	m := make(map[string]string)
	s := []int{1, 2, 3}
	s = append(s, 4)

	m["name"] = "Wilmer"
	fmt.Println(m)

	for i, v := range s {
		fmt.Println(i, " : ", v)
	}

	p := 34
	h := &p
	fmt.Println(p, h)
	fmt.Println(p, *h)

	c := make(chan string)
	go func() { // anonymous function
		fmt.Println("Starting function")
		time.Sleep(5 * time.Second)
		c <- "end channel"
	}()
	<-c
}
