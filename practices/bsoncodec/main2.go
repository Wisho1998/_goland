package main

//
//import "fmt"
//import "encoding/json"
//
//type Base struct {
//	Firstname string `json:"first"`
//}
//
//type A struct {
//	Base
//	Lastname string `json:"last"`
//}
//
//type B struct {
//	Base
//	Lastname string `json:"lastname"`
//}
//
//func main() {
//	john := A{Base: Base{Firstname: "John"}, Lastname: "Doe"}
//	john1 := B(john)
//	john_json, _ := json.Marshal(john)
//	john1_json, _ := json.Marshal(john1)
//	fmt.Println(string(john_json))
//	fmt.Println(string(john1_json))
//}
