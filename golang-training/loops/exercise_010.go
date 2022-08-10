package main

import "fmt"

func main() {

	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print2")
	case (3 == 4):
		fmt.Println("prints")
	case (4 == 5):
		fmt.Println("also true,does it print?")
	default:
		fmt.Println("default case")

	}

}
