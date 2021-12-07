package main

import . "fmt"

func Sum(x, y int) int {
	return x + y
}

func GetMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	commands := []struct {
		summary string
		command string
	}{
		{"initialize test", "go test"},
		{"show coverage", "go test -cover"},
		{"create file test", "go test -coverprofile=coverage.out"},
		// coverage
		{"show summary coverage", "go tool cover -func=coverage.out"},
		{"show summary coverage pretty", "go tool cover -html=coverage.out"},
		// profiling
		{"create file test cpu", "go test -cpuprofile=cpu.out"},
		{"enter detail cpu test", "go tool pprof cpu.out"},
		{"show detail cpu test", "top"},
		{"list detail cpu method by line", "list Fibonacci"},
	}
	for _, v := range commands {
		Println("Summary: ", v.summary, "\tCommand: ", v.command)
	}
}
