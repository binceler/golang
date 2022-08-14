package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {

	sa1 := secretAgent{
		person: person{
			first: "busra,",
			last:  "inc",
			age:   123,
		},
		ltk: true,
	}

	p1 := person{
		first: "james",
		last:  "bond",
		age:   40,
	}

	p2 := person{
		first: "miss",
		last:  "moneypenny",
		age:   32,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)

	fmt.Println(sa1)
	fmt.Println(sa1.person)
	fmt.Println(sa1.person.first)
	fmt.Println(sa1.ltk)

}
