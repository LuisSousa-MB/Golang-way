package model

import "fmt"

type Windows struct{}

func (w *Windows) InserirNaPortaUSB() {
	fmt.Println("O conector USB foi acoplado no computador Windows.")
}
