package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	x = append(x[:3], x[6:]...)
	fmt.Println(x)

	xi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println("--------------------")

	xi = append(xi[:4], xi[5:]...)
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println("--------------------")
}
