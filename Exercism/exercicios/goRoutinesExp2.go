package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var contador int64
var mu sync.Mutex

const x = 100000

func main() {
	fmt.Println("CPU´S: ", runtime.NumCPU())
	CriarGoRountines(x)
	wg.Wait()
	fmt.Println("Total de Go Routines: \t", x, "\nResultado do contador: \t", contador)
}

func CriarGoRountines(x int) {
	wg.Add(x)
	for i := 0; i < x; i++ {
		go func() {
			mu.Lock()
			v := contador
			runtime.Gosched()

			v++

			contador = v
			//fmt.Println("Contando: ", contador)
			mu.Unlock()
			wg.Done()
		}()
	}
}
