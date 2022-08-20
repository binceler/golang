package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6}
	x := boo(xi...)

	fmt.Println(x)
}
func boo(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for _, v := range x {
		sum += v
	}

	return sum
}
