package main

import "fmt"

func main() {

	x := make([]int, 10, 100)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[0] = 32
	x[4] = 42

	fmt.Println(x)

	x = append(x, 322)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}
