package main

import (
	"fmt"

	"github.com/jakoblorz/annotate"
)

type Person struct {
	Name string `default:"John"`
	Age  int    `default:"42"`
}

func (p Person) Default() *Person {
	p.Name = "Jane"
	return &p
}

type ReferencedPerson struct {
	Person annotate.DX[Person]
}

func main() {
	p := ReferencedPerson{}

	fmt.Printf("%+v\n", p.Person.Value()) // &{Name:Jane Age:42}
}
