package main

import "fmt"

func main() {
	f := foo()
	fmt.Println(f)
	a, b := bar()
	fmt.Println(a, b)
}

func foo() int {
	return 123
}

func bar() (int, string) {

	a := 456
	b := "456"

	return a, b

}
