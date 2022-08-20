package main

import "fmt"

func main() {
	defer foo() //defer ertelemek için kullanılır. foo yu main fonksiyonunda en sonda çalıştıracak
	bar()

}

func foo() {
	fmt.Println("hello world")
}

func bar() {
	fmt.Println("abcdef")
}
