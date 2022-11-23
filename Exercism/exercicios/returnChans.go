package main

import "fmt"

func main() {
	c := incrementor()
	cSum := puller(c)
	for n := range cSum {
		fmt.Println(n)
	}
}

func incrementor() chan int { // Cria um canal OUT e envia valores de 0 - 9.
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c chan int) chan int { // Recebe valores de um canal e retorna a soma dos valores recebidos.
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
			fmt.Println("Somando ", sum, " + ", n)
		}
		out <- sum
		close(out)
	}()
	return out
}
