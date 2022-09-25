package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var arr [5]int
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println(arr)
	fmt.Println(arr[2:4])
	m := arr[2:5]
	m[0] = 11
	fmt.Println(arr)
	var sl []int
	fmt.Println(sl)
	fmt.Println(len(sl), cap(sl))
	var krr [5]int = [5]int{1, 2, 3, 4, 5}
	for _, value := range krr {
		fmt.Println(value)
	}
}
