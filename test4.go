package main

import "fmt"

func main() {
	sum := 0
	var arr [5]int
	for i := 0; i < 5; i++ {
		fmt.Println("Enter the numbers")
		fmt.Scanf("%d\n", &arr[i])
	}
	for i := 0; i < 5; i++ {
		sum += arr[i]
	}
	fmt.Println("The sum of the numbers is:")
	fmt.Println(sum)
	var srr [5]string
	for i := 0; i < 5; i++ {
		fmt.Println("Enter the string")
		fmt.Scanf("%s\n", &srr[i])
	}
	for index, value := range srr {
		fmt.Println(index, value)
	}

	var nums [5]float64
	for i := 0; i < 5; i++ {
		fmt.Println("Enter the numbers")
		fmt.Scanf("%f\n", &nums[i])
	}

	for index, value := range nums {
		fmt.Println(index, value)
	}
}
