package main

import "fmt"

func main() {

	/*
		for init; condition; post {
			fmt.Println("Hello World!")
		}
	*/

	for i := 0; i <= 100; i++ {
		fmt.Println("Hello World!", i)
	}
}
