package main

import "fmt"

func main() {
	for x := 0; x < 20; x++ {
		if x%2 == 0 {
			continue
		}
		fmt.Println("counting odd numbers: ", x)
	}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println("Inner")
		}
		fmt.Println("Outer")
	}
}
