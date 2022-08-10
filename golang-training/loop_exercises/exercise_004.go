package main

import "fmt"

func main() {
	bd := 1992

	for {
		if bd == 2023 {
			break
		}
		fmt.Println(bd)
		bd++
	}

}
