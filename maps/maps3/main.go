package main

import "fmt"

func main() {
	am := map[string]int{
		"Drew":      28,
		"Aubri":     29,
		"AsherRose": 3,
	}
	fmt.Println(am)

	delete(am, "Drew")

	fmt.Println(am)
}
