package main

import "fmt"

func main() {

	f := func() {
		for i := 1; i < 3; i++ {
			fmt.Println(i)
		}
	}
	f()
	fmt.Printf("%T\n", f)

}
