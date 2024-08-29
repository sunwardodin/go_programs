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

	// Nested loop
	for i := 0; i < 5; i++ {
		fmt.Println("--")
		for j := 0; j < 5; j++ {
			fmt.Printf("outer loop %v \t inner loop %v\n", i, j)
		}
	}
	fmt.Println()

	// For range loops
	xi := []int{42, 43, 44, 45, 46, 47}
	for i, v := range xi {
		fmt.Println("ranging over a slice", i, v)
	}
	fmt.Println()

	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	for k, v := range m {
		fmt.Println("ranging over a map", k, v)
	}
}
