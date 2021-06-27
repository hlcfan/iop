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

	// fmt.Println("===Puts")
	// pp.Puts(alex)
	// fmt.Println("===Inspect")
	pp.Inspect(alex)

	bob := Person{
		ID:    2,
		Name:  "bob",
		Phone: "9876688",
	}

	people := []Person{alex, bob}
	fmt.Println("===Inspect Slice")
	pp.Inspect(people)

	fmt.Println("===Inspect Int")
	pp.Inspect(1)

	fmt.Println("===Inspect Float")
	pp.Inspect(1.12345678)
}
