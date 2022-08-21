package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("Foo ran")
}
func bar() {
	fmt.Println("bar ran")
}
