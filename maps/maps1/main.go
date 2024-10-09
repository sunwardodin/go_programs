package main

import "fmt"

func main() {
	am := map[string]int{
		"Drew":      28,
		"Aubri":     29,
		"AsherRose": 3,
	}
	fmt.Println("The age of Asher and Rose is ", am["AsherRose"])

	an := make(map[string]int)
	an["Nevada"] = 2
	an["Utah"] = 1
	an["Idaho"] = 5
	fmt.Println(an)
	fmt.Printf("%#v\n", an)
	fmt.Println(len(an))
}
