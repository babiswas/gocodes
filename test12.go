package main

import (
	"fmt"
	"sync"
)

func odd_gen(odd chan<- int, evn <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < 100; {
		odd <- i
		fmt.Println(<-evn)
		i = i + 2
	}
	close(odd)
}

func even_chan(odd <-chan int, evn chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 100; {
		fmt.Println(<-odd)
		evn <- i
		i = i + 2
	}
	close(evn)
}

func main() {
	var wg sync.WaitGroup
	odd := make(chan int)
	evn := make(chan int)
	wg.Add(1)
	go odd_gen(odd, evn, &wg)
	wg.Add(1)
	go even_chan(odd, evn, &wg)
	wg.Wait()
}
