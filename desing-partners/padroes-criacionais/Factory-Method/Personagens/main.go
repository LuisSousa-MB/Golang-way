package main

import (
	"Personagens-RPG/controller"
	"Personagens-RPG/database"
	"fmt"
)

func main() {
	database.ConectaComDB()
	fmt.Println("Gerador de personagem")
	novoPersongem := controller.GeradorDePersonagem()
	controller.ExibePersonagem(*novoPersongem)

}
