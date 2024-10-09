package main

import "fmt"

func main() {
	am := map[string]int{
		"Drew":      28,
		"Aubri":     29,
		"AsherRose": 3,
	}
	fmt.Println(am)

	v, ok := am["Georgey"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("Key does not exist.")
	}

	if v, ok := am["Drew"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("Key does not exist.")
	}

	if b, ok := am["Hag"]; !ok {
		fmt.Println(b)
	} else {
		fmt.Println("Key does not exist.")
	}
}
