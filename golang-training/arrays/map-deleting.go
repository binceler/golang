package main

import "fmt"

func main() {

	m := map[string]int{
		"james": 43,
		"busra": 32,
	}
	fmt.Println(m)

	delete(m, "james")

	fmt.Println(m)

}
