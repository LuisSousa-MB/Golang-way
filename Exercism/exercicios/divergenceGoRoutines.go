package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	canal1 := make(chan int)
	canal2 := make(chan int)
	const funcoes = 5

	go envia(100, canal1)
	go direcionadora(funcoes, canal1, canal2)

	for v := range canal2 {
		fmt.Println(v)
	}
}

func envia(x int, c1 chan int) {
	for i := 0; i < x; i++ {
		c1 <- i
	}
	close(c1)
}
func direcionadora(f int, c1, c2 chan int) {
	var wg sync.WaitGroup

	for i := 0; i < f; i++ {
		wg.Add(1)
		go func() {
			for v := range c1 {
				c2 <- trabalho(v)
			}
			wg.Done()
		}()

	}
	wg.Wait()
	close(c2)
}
func trabalho(n int) int {
	time.Sleep(time.Second) //* time.Duration(rand.Intn(2)))
	return n
}
