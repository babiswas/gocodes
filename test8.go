package main

import "fmt"

func main() {
	var queue []int
	queue = append(queue, 1)
	queue = append(queue, 2)
	queue = append(queue, 3)
	queue = append(queue, 4)

	for len(queue) > 0 {
		fmt.Println(queue[0])
		queue = queue[1:]
	}
}
