package main

import (
	"fmt"

	"github.com/sunwardodin/puppy"
)

func main() {
	puppy.From13()

	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println(puppy.BigBark())
	fmt.Println(puppy.BigBarks())

}
