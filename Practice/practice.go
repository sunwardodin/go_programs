package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	/*
		s, i, f := "Happiness", 42, 42.42
		fmt.Printf("%v is of type %T\n", i, i)
		fmt.Printf("%v is of type %T\n", f, f)
		fmt.Printf("%v is of type %T\n", s, s)

		x, y, z := 747, 911, 90210
		fmt.Printf("%d \t\t %b \t\t %#X\n",x,x,x)
		fmt.Printf("%d \t\t %b \t\t %#X\n",y,y,y)
		fmt.Printf("%d \t\t %b \t %#X\n",z,z,z)
	*/

	var x uint8 = 255
	var y int8 = 127

	fmt.Printf("%v is of type %T\n", x, x)
	fmt.Printf("%v is of type %T\n", y, y)

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)
}
