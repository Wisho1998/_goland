package main

import (
	"fmt"
	"protoOptions/pb"
)

func main() {
	newUser := &pb.User{
		Id:   "",
		Name: "",
	}
	fmt.Println("hola", newUser.ProtoReflect().Descriptor().Parent())
}
