package main

import "fmt"

func main() {
	fmt.Println("Teste")

	for i := 0; i <= 50; i++ {
		if i%5 == 0 {
			fmt.Println(i, " é um número divisível por 5")
		}
	}
}
