package main

import "fmt"

func test_func() (result int) {
	result = 4 * 2
	return
}

func main() {
	num := test_func()
	fmt.Println(num)
}
