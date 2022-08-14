package main

import "fmt"

func main() {

	x := struct {
		first_name string
		last_name  string
		age        int
		marriage   bool
	}{
		first_name: "busra",
		last_name:  "inc",
		age:        30,
		marriage:   false,
	}

	fmt.Println(x, x.first_name, x.last_name, x.age, x.marriage)

}
