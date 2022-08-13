package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Chocolate", "martini"}
	fmt.Println(jb)
	mp := []string{"miss", "moneypenny", "strawberry", "hazelnut"}
	fmt.Println(mp)

	xp := [][]string{jb, mp} //cok boyutlu dilim
	fmt.Println(xp)

}
