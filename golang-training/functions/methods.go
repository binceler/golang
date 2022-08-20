package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// func (r receiver ) identifier(parameters) (return(S)) { code }
func (s secretAgent) speak() {
	fmt.Println("I am: ", s.first)
}

func main() {

	sa1 := secretAgent{
		person: person{
			"james",
			"bond",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			"money",
			"penny",
		},
		ltk: false,
	}

	fmt.Println("hello, playground", sa1)

	sa1.speak()
	sa2.speak()

}
