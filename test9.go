package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num []int
	for i := 0; i < 5; i++ {
		num = append(num, rand.Intn(10))
	}
	fmt.Println(num)
	test := make([]int, len(num))
	copy(test, num)
	for _, value := range test {
		fmt.Println(value)
	}
}
