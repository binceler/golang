package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	fmt.Println(x[1])
	fmt.Println(x[:])
	fmt.Println(x[2:])
	fmt.Println(x[1:4])

	for i, v := range x {
		fmt.Println(i, v)
	}

	for i := 1; i <= 4; i++ {
		fmt.Println(i, x[i])
	}
}
