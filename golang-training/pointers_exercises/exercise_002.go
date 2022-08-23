package main

import "fmt"

type person struct {
	first string
}

func main() {
	p1 := person{
		first: "abc",
	}

	changeMe(&p1)
	fmt.Println(p1)
	fmt.Println(&p1.first)
}
func changeMe(p *person) {
	p.first = "moneypenny"
	// (*p).first = "moneypenny"
}
