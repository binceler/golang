package main

import "fmt"

func main() {

	foo()
	bar(map[string]int{"abc": 1})
	s1 := woo("moneypenny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")

	fmt.Println(x, y)

}

//func (r receiver) identifier(parameters) (return(s)) {....code...}

func foo() {
	fmt.Println("hello world from foo")
}

//everything in GO is pass by value
func bar(s map[string]int) {
	fmt.Println("hello", s)
}

func woo(s string) string {
	return fmt.Sprint("hello", s)
}

func mouse(fn, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, `, says "Hello"`)
	b := false
	return a, b
}
