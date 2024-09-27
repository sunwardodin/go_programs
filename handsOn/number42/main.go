package main

import (
	"fmt"
)

func main() {
	a := []int{0, 1, 2, 3, 4}
	for i := 0; i < 5; i++ {
		fmt.Printf("%T %v \n", a[i], a[i])
	}

	bSlice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for k := range bSlice {
		fmt.Println(k, bSlice[k])
	}

	fmt.Println(bSlice[:5])
	fmt.Println(bSlice[5:])
	fmt.Println(bSlice[2:7])
	fmt.Println(bSlice[1:6])
}
