package main

import (
	"fmt"

	"github.com/hlcfan/iop"
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
	// iop.Puts(alex)
	// fmt.Println("===Inspect")
	iop.Inspect(alex)

	bob := Person{
		ID:    2,
		Name:  "bob",
		Phone: "9876688",
	}

	people := []Person{alex, bob}
	fmt.Println("===Inspect Slice")
	iop.Inspect(people)

	fmt.Println("===Inspect Int")
	iop.Inspect(1)

	fmt.Println("===Inspect Float")
	iop.Inspect(1.12345678)
}
