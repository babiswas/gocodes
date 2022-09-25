package main

import "fmt"

func main() {
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	cp := make([]int, len(arr))
	copy(cp, arr[:5])
	fmt.Println(cp)
	for _, value := range cp {
		fmt.Println(value)
	}
}
