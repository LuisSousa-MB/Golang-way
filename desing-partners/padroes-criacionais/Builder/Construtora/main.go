package main

import (
	"Construtora/models"
	"fmt"
)

func main() {
	construtorNormal := models.GetConstrutor("Normal")
	construtorModerno := models.GetConstrutor("Moderno")

	diretor := models.NovoDiretor(construtorNormal)
	construcaoNormal := diretor.Construir()

	fmt.Printf("Tipo de porta da construção: %s\n", construcaoNormal.TipoDePorta)
	fmt.Printf("Tipo de janelas da construção: %s\n", construcaoNormal.TipoDeJanela)
	fmt.Printf("Tipo de piso da construção: %s\n", construcaoNormal.TipoPiso)
	fmt.Printf("Quantidade de janelas: %d\n", construcaoNormal.QuantidadeDeJanelas)
	fmt.Printf("Quantidade de andares: %d\n", construcaoNormal.QuantidadeDeAndares)
	//fmt.Printf("Chaminé:\n", construcaoNormal.Chamine)

	diretor.SetConstrutor(construtorModerno)
	construcaoModerna := diretor.Construir()
	fmt.Println()
	fmt.Printf("Tipo de porta da construção: %s\n", construcaoModerna.TipoDePorta)
	fmt.Printf("Tipo de janelas da construção: %s\n", construcaoModerna.TipoDeJanela)
	fmt.Printf("Tipo de piso da construção: %s\n", construcaoModerna.TipoPiso)
	fmt.Printf("Quantidade de janelas: %d\n", construcaoModerna.QuantidadeDeJanelas)
	fmt.Printf("Quantidade de andares: %d\n", construcaoModerna.QuantidadeDeAndares)
	//fmt.Printf("Chaminé:\n", construcaoModerna.Chamine)
}
