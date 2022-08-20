package main

import "fmt"

func main() {
	foo()
	x := boo(1, 2, 3)

	fmt.Println(x)
}

func foo() {
	fmt.Println("merhaba")
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

//func (r receiver) identifier (parameter(s)) (retrun(s)) { code }
