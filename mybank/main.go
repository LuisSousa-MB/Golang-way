package main

import (
	"fmt"
	"mybank/database"
	"mybank/routes"
	"os"
	"time"
)

var saque float64
var deposito float64
var transferencia float64

func main() {
	database.ConectaComDB()
	routes.HandleRequests()
}
func standby() {

	for {

		fmt.Println("\n1 - Acessar uma conta\n2 - Cadastrar nova conta")

		var opcao int

		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			autenticado := Autenticar()
			fmt.Println("Acesso", autenticado)
			if autenticado {
				fmt.Println("Bem vindo ao seu banco!")
				for {
					ExibeMenu()
					comando := LerComando()

					switch comando {
					case 1:
						fmt.Println("Exibir saldo")
					case 2:
						fmt.Println("\nInforme o valor que deseja sacar:")
						fmt.Scan(&saque)
					//	contaDaSilvia.Sacar(saque)
					case 3:
						fmt.Println("Informe o valor que deseja depositar:")
						fmt.Scan(&deposito)
					//	contaDaSilvia.Depositar(deposito)
					case 4:
						fmt.Println("Informe o valor que deseja tranferir:")
						fmt.Scan(&transferencia)
						//contaDaSilvia.Transferir(transferencia, &contaDoJose)
					case 0:
						fmt.Println("Saindo do programa...")
						time.Sleep(1 * time.Second)
						os.Exit(0)

					}
				}

			} else {
				fmt.Println("Acesso negado")
			}

		case 2:
			for {
				var opcao int
				fmt.Println("Tipo de conta: \n\n1 - Conta corrente\n2 - Conta poupança")
				fmt.Scan(&opcao)

				if opcao == 1 {
					/* novaConta := models.ContaCorrente{}
					novaConta.CadNovaConta() */
					break
				} else if opcao == 2 {
					/* novaConta := models.ContaPoupanca{}
					novaConta.CadNovaConta() */
					break
				} else {
					fmt.Println("Opção inválida")
					fmt.Println()
				}

			}

		case 0:
			fmt.Println("Saindo do programa...")
			time.Sleep(1 * time.Second)
			os.Exit(0)

		}

	}
}
func ExibeMenu() {
	fmt.Println("\n\nSelecione uma opção:")
	fmt.Println("\n1 - saldo")
	fmt.Println("2 - Sacar")
	fmt.Println("3 - Depositar")
	fmt.Println("4 - Transferir")
	fmt.Println("0 - Sair")

}
func LerComando() int {
	var comando int
	fmt.Scan(&comando)
	return comando
}
func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	if conta.Sacar(valorDoBoleto) {
		fmt.Println("Boleto pago no valor de: R$", valorDoBoleto)
	} else {
		fmt.Println("Saldo insuficiente para o pagamento")

	}
}
func Autenticar() bool {
	return true
}

type verificarConta interface {
	Sacar(valor float64) bool
}
