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

	fmt.Println(p1, p1.first, p1.last, p1.favFlavors)
	fmt.Println(p2, p2.first, p2.last, p2.favFlavors)

}
