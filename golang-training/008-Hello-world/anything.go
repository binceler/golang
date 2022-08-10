package main

import "fmt"

func main() {
	fmt.Println("Hello everyone")
	foo()
	fmt.Println("something wrong")

	for i := 0; i < 100; i++ {
		if i%13 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}

func foo() {
	fmt.Println("I'm in foo")

}

func bar() {
	fmt.Println("and then we exited!")

}

// control flow:
// (1) sequence
// (2) loop; iterative
// (3) conditional
