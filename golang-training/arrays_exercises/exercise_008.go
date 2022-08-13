package main

import "fmt"

func main() {
	m := map[string][]string{
		"james":      []string{`shaken`, `martinis`, `women`},
		"moneypenny": []string{`james`, `literature`, `computer science`},
		"no_dr":      []string{`being evil`, `ice cream`, `sunsets`},
	}
	fmt.Println(m)
	for _, v := range m {
		fmt.Println(v)
	}

}
