package main

import "fmt"

var x int

func main() {
	fmt.Println(x)
	fmt.Println(x)
	foo()
	{
		y := 42
		fmt.Println(y)
	}
	fmt.Println(x)

}

func foo() {
	fmt.Println("hello")
	fmt.Println(x)
	x++
	fmt.Println(x)
}
