package main

import (
	"Singleton/model"
	"fmt"
)

func main() {

	for i := 0; i < 30; i++ {
		go model.GetInstance()
	}
	go model.GetInstance()
	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}
