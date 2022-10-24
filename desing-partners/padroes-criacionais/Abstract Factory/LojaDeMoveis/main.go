package main

import (
	"LojaDeMoveis/models"
	"fmt"
)

func main() {
	fmt.Println("Móveis gerados")
	fmt.Println()
	factory, _ := models.GetMaterialStyle("Art")
	artCadeira := factory.CriarCadeira()

	factory, _ = models.GetMaterialStyle("Vitoriano")
	vitorianoCadeira := factory.CriarCadeira()
	vitorianoArmario := factory.CriarArmario()

	factory, _ = models.GetMaterialStyle("moderno")
	modernoCama := factory.CriarCama()
	modernoPoltrona := factory.CriarPoltrona()

	ExibeDetalhesCadeira(artCadeira)
	ExibeDetalhesCadeira(vitorianoCadeira)
	ExibeDetalhesCadeira(vitorianoArmario)
	ExibeDetalhesCama(modernoCama)
	ExibeDetalhesPoltrona(modernoPoltrona)
}

func ExibeDetalhesArmario(c models.IArmario) {
	fmt.Printf("Descrição: %s", c.GetDescricao())
	fmt.Println()
	fmt.Printf("Estilo: %s", c.GetStyle())
	fmt.Println()
}
func ExibeDetalhesCadeira(c models.ICadeira) {
	fmt.Printf("Descrição: %s", c.GetDescricao())
	fmt.Println()
	fmt.Printf("Estilo: %s", c.GetStyle())
	fmt.Println()
}
func ExibeDetalhesCama(c models.ICama) {
	fmt.Printf("Descrição: %s", c.GetDescricao())
	fmt.Println()
	fmt.Printf("Estilo: %s", c.GetStyle())
	fmt.Println()
}
func ExibeDetalhesMesa(c models.IMesa) {
	fmt.Printf("Descrição: %s", c.GetDescricao())
	fmt.Println()
	fmt.Printf("Estilo: %s", c.GetStyle())
	fmt.Println()
}
func ExibeDetalhesPoltrona(c models.IPoltrona) {
	fmt.Printf("Descrição: %s", c.GetDescricao())
	fmt.Println()
	fmt.Printf("Estilo: %s", c.GetStyle())
	fmt.Println()
}
