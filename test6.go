package main

import (
	"errors"
	"fmt"
	"os"
)

type Stack struct {
	nums  []int
	index int
	size  int
}

func (s *Stack) init(n int) {
	s.nums = make([]int, n)
	s.index = -1
	s.size = n
}

func (s *Stack) push(num int) error {
	s.index += 1
	if s.index == s.size {
		return errors.New("Overflow")
	}
	s.nums[s.index] = num
	return nil
}

func (s *Stack) pop() (int, error) {
	if s.index == -1 {
		return -1, errors.New("Empty")
	}
	num := s.nums[s.index]
	s.nums[s.index] = 0
	s.index -= 1
	return num, nil
}

func handle(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	s := Stack{}
	s.init(4)
	err := s.push(3)
	handle(err)
	err = s.push(4)
	handle(err)
	err = s.push(5)
	handle(err)
	err = s.push(6)
	handle(err)
	fmt.Println(s)
	fmt.Println("Process the stack for elements:")
	num, err := s.pop()
	handle(err)
	fmt.Println(num, s)
	num, err = s.pop()
	handle(err)
	fmt.Println(num, s)
	num, err = s.pop()
	handle(err)
	fmt.Println(num, s)
	num, err = s.pop()
	handle(err)
	fmt.Println(num, s)
	num, err = s.pop()
	handle(err)
	fmt.Println(num, s)
}
