package main

import "fmt"

func main() {
	num := 0
	var stack []int
	stack = append(stack, 1)
	stack = append(stack, 2)
	stack = append(stack, 3)
	stack = append(stack, 4)
	for len(stack) > 0 {
		num = stack[len(stack)-1]
		fmt.Println(num)
		stack = stack[:len(stack)-1]
	}
}
