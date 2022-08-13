package main

import "fmt"

func main() {

	m := map[string]int{
		"James":      32,
		"moneypenny": 21,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["barnabas"]) //cıktısı 0(sıfır) olur cünkü mapte tanımlı değil

	v, ok := m["barnabas"]
	fmt.Println(v)
	fmt.Println(ok)

	if z, okey := m["James"]; okey {
		fmt.Println(m["busra"])
		fmt.Println(z)
	}

}
