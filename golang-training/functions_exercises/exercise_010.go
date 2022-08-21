package main

import "fmt"

func main() {
	f := foo()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	g := foo()
	fmt.Println(g())
}

func foo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
