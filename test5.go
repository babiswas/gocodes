package main

import (
	"fmt"
	"os"
)

func find_type(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Integer")
	case string:
		fmt.Println("String")
	case float64:
		fmt.Println("Float64")
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	find_type("hello")
	find_type(21)
	find_type(2.33)
	for index, value := range os.Args[1:] {
		fmt.Println(index, value)
	}
}
