package controllers

import (
	"Xburguer/models"
	"fmt"
)

func GerarPedido() {
	var preparadorSelecionado models.IPreparador
	for x := 1; x > 0; x -= 1 {
		var pedido int
		fmt.Scan(&pedido)
		switch pedido {
		case 1:
			preparadorSelecionado = models.GetPreparador("ComboX")
		case 2:
			preparadorSelecionado = models.GetPreparador("ComboY")
		case 3:
			preparadorSelecionado = models.GetPreparador("ComboZ")
		case 4:
			preparadorSelecionado = models.GetPreparador("ComboS")
		case 5:
			preparadorSelecionado = models.GetPreparador("ComboMonsterX")
		case 6:
			preparadorSelecionado = models.GetPreparador("IceCloudX")
		case 7:
			preparadorSelecionado = models.GetPreparador("OnlyFries")
		default:
			fmt.Println("Pedido inválido...\nTente novamente:")
			x += 1
		}
	}
	gerente := models.NovoGerente(preparadorSelecionado)
	comboSelecionado := gerente.Preparar()
	fmt.Printf("Pedido pronto! \n\n%s", comboSelecionado.Nome+"\n\n	"+comboSelecionado.Descricao)
	fmt.Println()
}
func ExibirMenu() {
	fmt.Println("Bem vindo a Xburguer!\n\nQual o seu pedido:")
	fmt.Println()
	fmt.Println("\n1 - Combo X...........(Hamburguer X, Refrigerante, Batatas-médias, Sobremesa do dia, Brinde")
	fmt.Println("\n2 - Combo Y...........(Hamburguer Y, Refrigerante, Batatas-grande, Sobremesa do dia")
	fmt.Println("\n3 - Combo Z...........(Hamburguer Z, Suco natural, Batatas-médias, Brinde")
	fmt.Println("\n4 - Combo S...........(Hamburguer S, Suco natural, Salada X, Sobremesa light")
	fmt.Println("\n5 - Combo MonsterX....(Hamburguer Triple X, Refrigerante-XBig, Batatas-Grandes, IceCouldx, brinde")
	fmt.Println("\n6 - IceCould X........(Sorvete com flocos, Chocolate meio amargo e Frutas vermelhas")
	fmt.Println("\n7 - OnlyFries.........(Batatas-médias")
	fmt.Println()
}
