package main

import (
	"fmt"

	"github.com/hlcfan/pp"
)

type Person struct {
	ID    int
	Name  string
	Phone string
}

func main() {
	alex := Person{
		ID:    1,
		Name:  "alex",
		Phone: "12345678",
	}

	pp.Puts(alex)

	bob := Person{
		ID:    2,
		Name:  "bob",
		Phone: "9876688",
	}

	people := []Person{alex, bob}
	pp.Puts(people)

	pp.Puts(1)

	fmt.Println("===Inspect Float")
	pp.Puts(1.12345678)
}
