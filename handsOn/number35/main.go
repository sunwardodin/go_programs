package main

import (
	"fmt"
	"math/rand"
)

func main() {
	xi := []int{42, 43, 44, 45, 46, 47}

	for i, v := range xi {
		fmt.Printf("%d, %d\n", i, v)
	}

	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	for i, v := range m {
		fmt.Printf("%s, %d\n", i, v)
	}
	age := m["James"]
	fmt.Println(age)

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

	for i := 0; i <= 100; i++ {
		if z := rand.Intn(5); z == 3 {
			fmt.Printf("x is 3\n")
		}
	}
}
