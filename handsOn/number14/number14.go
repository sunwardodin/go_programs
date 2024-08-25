package main

import "fmt"

var stuff int = 3

const stuff2 string = "Stuff 2"

func aNumber() string {
	return "This is a function for the number 10"
}

func main() {
	s1 := "hello"

	fmt.Println(stuff)
	fmt.Println(stuff2)
	fmt.Println(aNumber())
	fmt.Printf("the value of stuff is %v and the type of stuff is %T \n", stuff, stuff)
	fmt.Printf("the value of stuff2 is %v and the type of stuff2 is %T \n", stuff2, stuff2)
	fmt.Printf("the value of s1 is %v and the type of s1 is %T \n", s1, s1)
}
