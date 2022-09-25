package main

import (
	"flag"
	"fmt"
)

func main() {
	comm := flag.String("comm", "hello", "A srting value for comm")
	nums := flag.Int("val", 12, "An int value for num")
	test := flag.Bool("test", false, "A boolean value")
	var tes string
	flag.StringVar(&tes, "tes", "hello", "A string variable")
	flag.Parse()
	fmt.Println("The value of comm is:", *comm)
	fmt.Println("The value of nums is:", *nums)
	fmt.Println("The value of test is:", *test)
	fmt.Println("The value of tes is:", tes)
}
