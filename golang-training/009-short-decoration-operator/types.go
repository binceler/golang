package main

import "fmt"

var y = 42
var z = "Shaken , not stirred to me!"

var a string = "bla bla bla"

var b string = `james said that,"i was so squared"`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
