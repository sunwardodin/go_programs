package main

import "fmt"

func main() {
	x1 := []string{"James", "Bond", "Shaken, not stirred"}
	x2 := []string{"Miss", "Moneypenny", "I'm 008."}
	x40 := [][]string{x1, x2}

	//fmt.Println(x40)

	for k, v := range x40 {
		fmt.Println(k)
		for i := 0; i < len(v); i++ {
			fmt.Println(v[i])
		}
	}
}
