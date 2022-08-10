package main

import "fmt"

func main() {

	if true {
		fmt.Println("one")
	}

	if false {
		fmt.Println("two")
	}

	if !true {
		fmt.Println("three")
	}

	if !false {
		fmt.Println("four")
	}

	if 2 == 2 {
		fmt.Println("five")
	}
	if 2 != 2 {
		fmt.Println("six")
	}
	if !(2 == 2) {
		fmt.Println("seven")
	}

	if !(2 != 2) {
		fmt.Println("eight")
	}

}
