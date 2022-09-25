package main

import "fmt"

func main() {
	test := struct {
		name string
		id   int
	}{"hello", 12}

	fmt.Println(test.name)
	fmt.Println(test.id)
	type Person struct {
		name string
	}
	ch := make(chan interface{}, 3)
	person := []Person{{name: "Bapan"}, {name: "hello"}, {name: "bello"}}
	for i := 0; i < 3; i++ {
		ch <- person[i]
	}
	close(ch)
	for value := range ch {
		fmt.Println(value.(Person).name)
	}

	ch1 := make(chan struct{}, 5)
	for i := 0; i < 5; i++ {
		ch1 <- struct{}{}
	}
	close(ch1)
	for val := range ch1 {
		fmt.Println(val)
	}
}
