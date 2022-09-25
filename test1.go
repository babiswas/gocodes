package main

import (
	"fmt"
	"os"
)

func main() {
	for _, value := range os.Args[1:] {
		switch value {
		case "hello", "bello":
			fmt.Println("Either hello or bello is selected:")
		case "mello":
			fmt.Println("Mello is selected")
		default:
			fmt.Println("Default block is executed:")
		}
	}

	whoamI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("I am an integer:")
		case float64:
			fmt.Println("I am a floating point value:")
		case string:
			fmt.Println("I am a string:")
		default:
			fmt.Println("I am the default block:")
		}
	}

	whoamI(12)
	whoamI("hello")
	whoamI(6.3444)
	whoamI(123)
}
