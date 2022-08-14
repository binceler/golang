package main

import "fmt"

type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {

	p1 := person{
		first:      "Busra",
		last:       "inc",
		favFlavors: []string{"rose"},
	}

	p2 := person{
		first:      "james",
		last:       "bond",
		favFlavors: []string{"rose", "roses"},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(m)

	for i, v := range m {
		fmt.Println(i, v)
	}

}
