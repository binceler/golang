package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {

	s := fmt.Sprintf("%v\t%v\t%v\t", x, y, z)
	t := fmt.Sprintf("%T\t%T\t%T\t", x, y, z)
	fmt.Println(s)
	fmt.Println(t)

}
