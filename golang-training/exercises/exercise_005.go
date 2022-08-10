package main

import "fmt"

type a int

var x a

var y int

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Printf("%T\n", y)
	fmt.Println(y)

}
