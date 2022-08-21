package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 7, 8, 4}

	a := foo(ii...)
	b := bar(ii)
	fmt.Println(a)
	fmt.Println(b)
}

func foo(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}

	return total
}

func bar(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}

	return total
}
