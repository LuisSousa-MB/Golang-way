package main

import (
	"Personagens-RPG/controller"
	"fmt"
)

func main() {
	fmt.Println("Gerador de personagem")
	novoPersongem := controller.GeradorDePersonagem()
	controller.ExibePersonagem(*novoPersongem)

}
