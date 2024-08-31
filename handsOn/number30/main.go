package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 0; i <= 42; i++ {
		x := rand.Intn(5)
		switch x {
		case 0:
			fmt.Println("x equals 0\t")
		case 1:
			fmt.Println("x equals 1\t")
		case 2:
			fmt.Println("x equals 2\t")
		case 3:
			fmt.Println("x equals 3\t")
		case 4:
			fmt.Println("x equals 4\t")
		default:
			fmt.Println("x sucks")
		}
		fmt.Println(i)
	}

}
