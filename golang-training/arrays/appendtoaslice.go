package main

import "fmt"

func main() {

	x := []int{4, 5, 6, 7}
	fmt.Println(x)

	x = append(x, 77, 88, 99, 1024)

	fmt.Println(x)

	y := []int{21, 32, 43, 54}

	x = append(x, y...)

	fmt.Println(x)

}
