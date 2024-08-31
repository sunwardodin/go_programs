package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)
	fmt.Println(x)
	fmt.Println(y)

	switch x {
	case 3, 2, 1, 0:
		if y < 4 {
			fmt.Println("x and y are both less than 4")
		}
		fallthrough
	case 4, 5, 6:
		fmt.Println("x is greater than or equal to 4 and less than or equal to 6")
	case 7, 8, 9:
		if y > 6 {
			fmt.Println("x and y are both greater than 6")
		}
		fallthrough
	default:
		fmt.Println("none of the previous cases were met")
	}
	if y != 5 {
		fmt.Println("y is not 5")
	}
}
