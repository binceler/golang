package main

import "fmt"

func main() {
	x := [5]int{45, 46, 47, 48, 49}

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", x)

}
