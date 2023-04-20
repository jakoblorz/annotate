# annotate

annotate is a selection of golang types and functions that make it easier to automate different aspects of normally tedious tasks with the help of performant reflection and generics.

## Installation

```bash
go get github.com/jakoblorz/annotate
```

## Usage

### Default Values

There are many different ways to implement default values of fields in golang. annotate provides a simple way to do so by using struct tags. It uses a type cache to precompute the default values, so that the performance impact is minimal.

```go
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
```

Defaulting objects is also possible, with an API borrowed from the sql.NullX types. The default value is computed lazily, so that it is only computed if the value is actually used 
and the enclosed value is stored in a pointer, so that it can be nil. Upon defaulting, `annotate.ApplyDefaults` will be applied to the value, too, so that the enclosed's fields are also defaulted.


```go
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
```


If you want to customize the default value on the field level, without relying on the field's tags, you can use `annotate.DX` which requires you to implement the `annotate.Defaulter` interface on the enclosed struct.

```go
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
```