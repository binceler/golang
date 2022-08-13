package main

import "fmt"

func main() {

	x := []int{4, 5, 6, 7}
	fmt.Println(x)
	x = append(x, 77, 66)
	fmt.Println(x)
	x = append(x[:1], x[3:]...)
	fmt.Println(x)
}
