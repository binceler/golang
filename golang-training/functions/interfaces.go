package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println(s.first, "- secretAgent speak")
}

func (p person) speak() {
	fmt.Println(p.first, " ", p.last, "- person speak")
}

type human interface {
	speak()
}

func bar(h human) {
	fmt.Println("I called human", h)
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "james",
			last:  "bond",
		},
		ltk: false,
	}

	p1 := person{
		first: "busra",
		last:  "inc",
	}

	sa1.speak()

	fmt.Println(p1)

	p1.speak()
	bar(p1)
	bar(sa1)
}
