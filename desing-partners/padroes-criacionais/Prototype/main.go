package main

import (
	"Prototype/model"
	"Prototype/prototype"
	"fmt"
)

func main() {
	arquivo1 := &model.Arquivo{Nome: "Arquivo1"}
	arquivo2 := &model.Arquivo{Nome: "Arquivo2"}
	arquivo3 := &model.Arquivo{Nome: "Arquivo3"}

	pasta1 := &model.Pasta{
		Filho: []prototype.Inode{arquivo1},
		Nome:  "Pasta1",
	}

	pasta2 := &model.Pasta{
		Filho: []prototype.Inode{pasta1, arquivo2, arquivo3},
		Nome:  "Pasta2",
	}
	fmt.Println("\nExibindo hierarquia da Pasta2")
	pasta2.Exibir("  ")

	copiaDaPasta := pasta2.Clonar()
	fmt.Println("\nExibindo hierarquia da pasta copiada")
	copiaDaPasta.Exibir("  ")
}
