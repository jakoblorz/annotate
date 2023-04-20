package main

import (
	"fmt"
	"github.com/jakoblorz/annotate"
)

type Person struct {
	Name string `default:"John"`
	Age  int    `default:"42"`
}

func main() {
	p := Person{
		Name: "John",
	}

	annotate.ApplyDefaults(&p)

	fmt.Printf("%+v\n", p) // {Name:John Age:42}
}
