package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

var mu sync.Mutex

func main() {

	numbers := make(chan int)
	quit := make(chan bool)

	var jackPot bool
	//miners := int(runtime.NumCPU())

	unlockNumber := 777

	for i := 0; i < 900; i++ {
		go func(miner int) {
			jackPot = false

			for i := 0; i < 1; {
				mu.Lock()
				if jackPot == false {

					fmt.Println("Minerador:", miner)

					luckNumber := rand.Intn(1000)
					if luckNumber == unlockNumber {
						fmt.Println("Bingo!", luckNumber)
						quit <- true
						jackPot = true
						i++
					} else {
						numbers <- luckNumber

					}
					mu.Unlock()
					runtime.Gosched()

				} else {
					runtime.Gosched()

				}

			}

		}(i)
	}
	go func() {
		<-quit
		close(numbers)
	}()

	for n := range numbers {
		fmt.Println("...", n)
	}
}
