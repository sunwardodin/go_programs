package main

import "fmt"

type person struct {
	first    string
	last     string
	iceCream []string
}

func main() {
	p1 := person{
		first:    "Drew",
		last:     "Mikels",
		iceCream: []string{"Chocolate", "Vanilla"},
	}

	fmt.Println(p1.first, p1.last)
}
