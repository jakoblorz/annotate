package main

import (
	"fmt"

	"github.com/jakoblorz/annotate"
)

type Person struct {
	Name string `default:"John"`
	Age  int    `default:"42"`
}

type ReferencedPerson struct {
	Person annotate.D[Person]
}

func main() {
	p := ReferencedPerson{}

	fmt.Printf("%+v\n", p.Person.Value()) // &{Name:John Age:42}
}
