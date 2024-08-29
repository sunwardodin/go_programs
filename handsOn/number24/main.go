package main

import (
	"fmt"

	"math/rand"
)

func init() {
	fmt.Println("This is where my program is initialized")
}

func main() {
	x := rand.Intn(250)
	fmt.Println(x)
	y := rand.Intn(3)
	if y == 0 {
		fmt.Println(y)
	}

	if x <= 100 {
		fmt.Println("between 0 and 100")
	} else if x >= 101 && x <= 200 {
		fmt.Println("between 101 and 200")
	} else if x >= 201 {
		fmt.Println("between 201 and 250")
	}

	switch x {
	case 250:
		fmt.Println("x is 250")
	case 3:
		fmt.Println("x is 3")
	case 0:
		fmt.Println("x is 0")
	default:
		fmt.Println("x sucks")
	}
}
