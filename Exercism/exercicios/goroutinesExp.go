package main

import (
	"fmt"
	"time"
)

func Worker(workerId int, msg chan int) {
	for res := range msg {
		fmt.Println("Worker: ", workerId, " Msg: ", res)
		time.Sleep(time.Second)
	}
}

func main() {
		
	msg := make(chan int)
	for w := 1; w < 1000000; w++ {
		go Worker(w, msg)
	}

	for i := 0; i < 1000000000; i++ {

		msg <- i
	}
}
