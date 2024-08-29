package main

import "fmt"

func main() {
	y := 5
	// Loops
	for i := 0; i < 5; i++ {
		fmt.Printf("counting to five: %v \t first for loop\n", i)
	}
	// a while loop
	for y < 10 {
		fmt.Printf("y is %v \t\t\t second for loop\n", y)
		y++
	}
	// break takes you out of the loop
	for {
		fmt.Printf("y is %v \t\t third for loop\n", y)
		if y > 20 {
			break
		}
		y++
	}
	// continue takes it to the next iteration
	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println("counting even numbers: ", i)
	}
}
