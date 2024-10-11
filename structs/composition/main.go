package main

import "fmt"

type foo int

func main() {
	var a foo = 42

	fmt.Printf("%T\n", a)
}
