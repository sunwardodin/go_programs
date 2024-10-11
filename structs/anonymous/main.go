package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := struct { //This is an anonymous struct
		first string
		last  string
		age   int
	}{
		"Drew",
		"Mikels",
		28,
	}

	p2 := person{
		"Aubri",
		"Mikels",
		28,
	}

	fmt.Printf("%T\n", p1)
	fmt.Printf("%T\n", p2)
}
