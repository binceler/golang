package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("I am ", p.first, p.last, "and I am ", p.age, "years old.")
}

func main() {
	p1 := person{
		first: "abc",
		last:  "def",
		age:   31,
	}

	p1.speak()

}
