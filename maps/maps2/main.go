package main

import "fmt"

func main() {
	am := map[string]int{
		"Drew":      28,
		"Aubri":     29,
		"AsherRose": 3,
	}
	am["George"] = 34

	for k, v := range am {
		fmt.Println(k, v)
	}

	for _, v := range am {
		fmt.Println(v)
	}

	for k := range am {
		fmt.Println(k)
	}
}
