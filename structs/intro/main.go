package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {

	p1 := person{
		"Drew",
		"Mikels",
		28,
	}

	p2 := person{
		"Aubri",
		"Mikels",
		28,
	}

	fmt.Println(p1, p2)
	fmt.Printf("%T \t %#v\n", p1, p1)
	fmt.Println(p1.first, p2.age)
}
