package main

import "fmt"

var y string

func main() {
	fmt.Println("value of string y", y, "ending")
	fmt.Printf("%T\n", y)
	y = "bla bla bla"

	fmt.Println(y)
}
