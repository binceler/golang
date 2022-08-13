package main

import "fmt"

func main() {

	x := map[string]int{
		"james": 32,
	}
	fmt.Println(x)

	x["todd"] = 35

	fmt.Println(x)

	for k, v := range x {
		fmt.Println(k)
		fmt.Println(v)
	}

	xi := []int{4, 5, 6}

	for k, v := range xi {
		fmt.Println(k, v)
	}

}
