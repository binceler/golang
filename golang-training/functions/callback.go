package main

import "fmt"

func main() {
	ii := []int{2, 1, 4, 3, 5, 6}
	s := sum(ii...)
	fmt.Println(s)

	s2 := even(sum, ii...)
	fmt.Println(s2)
}
func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}

	return total

}

func even(f func(xi ...int) int, vi ...int) int {

	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}

	return f(yi...)
}
