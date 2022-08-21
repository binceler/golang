package main

import "fmt"

func main() {
	n := factorial(3)
	fmt.Println(n)
	l := loopFac(4)
	fmt.Println(l)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func loopFac(n int) int {
	var fac int = 1
	for i := 1; i <= n; i++ {
		fac *= i
	}
	return fac
}
