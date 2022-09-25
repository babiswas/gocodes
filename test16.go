package main

import (
	"fmt"
	"math/rand"
)

func main() {
	num := make([][]int, 3)
	fmt.Println(num)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			num[i] = append(num[i], j+rand.Intn(21))
		}
	}
	fmt.Println(num)
}
