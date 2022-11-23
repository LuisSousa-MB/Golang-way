package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// ch := make(chan int)
	ch := make(chan int, 3)

	wg.Add(1)
	go print(ch)

	for i := 0; i < 6; i++ {
		value := i * 10
		fmt.Printf("Write %02d\n", value)
		ch <- value
	}
	close(ch)
	fmt.Println("Wait...")
	wg.Wait()
}

func print(ch <-chan int) {
	for c := range ch {
		fmt.Printf("Print %02d (len: %d)\n", c, len(ch))
	}
	wg.Done()

}
