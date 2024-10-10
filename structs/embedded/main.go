package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk   bool
	first string
}

func main() {
	sa1 := secretAgent{
		person: person{
			"Drew",
			"Mikels",
			28,
		},
		first: "Joseph",
		ltk:   true,
	}

	fmt.Println(sa1)
	fmt.Printf("%T \t %#v\n", sa1, sa1)
	fmt.Println(sa1.first, sa1.person.first)
}
