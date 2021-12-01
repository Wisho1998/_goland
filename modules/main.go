package main

import (
	. "fmt"
	"github.com/donvito/hellomod"
)

func main() {
	Println("Hello world")

	hellomod.SayHello()

	commands := []struct {
		summary string
		command string
	}{
		{"initialize module", "go mod init github.com/username/module"},
		{"download dependency", "go get github.com/donvito/hellomod"},
		{"clean unused deps", "go mod tidy"},
		{"show list downloads", "go mod download -json"},
	}

	for _, v := range commands {
		Println("Summary: ", v.summary, "\tCommand: ", v.command)
	}
}
