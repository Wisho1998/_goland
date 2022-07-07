package main

import (
	"encoding/json"
	"fmt"
)

type S1 struct {
	Name   string
	Nested []*S2
}

type S2 struct {
	Name string
}

func main() {
	instance := &S1{
		Name:   "property 1",
		Nested: []*S2{{Name: "nested 1"}, {Name: "nested 2"}},
	}

	cloneMyStruct, _ := CloneMyStruct[[]*S2](&instance.Nested)

	instance.Nested[0].Name = "nuevo"

	for _, s2 := range instance.Nested {
		fmt.Println(s2)
	}
	for _, s2 := range cloneMyStruct {
		fmt.Println(s2)
	}
}

type cloneStruct interface {
	[]*S2
}

func CloneMyStruct[T cloneStruct](orig *T) (T, error) {
	origJSON, err := json.Marshal(orig)
	if err != nil {
		return nil, err
	}

	var clone T
	if err = json.Unmarshal(origJSON, &clone); err != nil {
		return nil, err
	}

	return clone, nil
}
