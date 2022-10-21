package contas

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	cliente "projects/mybank/clientes"
	"time"
)

type ContaCorrente struct {
	Titular                        cliente.Titular
	NumAgencia, NumConta, Operacao int
	saldo                          float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) bool {

	saque := valorDoSaque <= c.saldo && valorDoSaque > 0
	if saque {
		c.saldo -= valorDoSaque
		fmt.Println("\nSaque realizado no valor de R$:", valorDoSaque, "\nsaldo atual R$:", c.saldo, "\n__________________________________")
		time.Sleep(3 * time.Second)
		return true
	} else if valorDoSaque < 0 {
		fmt.Println("\nO valor solicitado é inválido...\n__________________________________")
		time.Sleep(3 * time.Second)
		return false
	} else {
		fmt.Println("\nsaldo insuficiente...\n__________________________________")
		time.Sleep(3 * time.Second)
		return false

	}

}
func (c *ContaCorrente) Depositar(valor float64) {
	deposito := valor > 0
	if deposito {
		c.saldo += valor
		fmt.Println("\nDepósito realizado no valor de R$:", valor, "\nsaldo atual R$:", c.saldo, "\n__________________________________")
		time.Sleep(3 * time.Second)
	} else if valor < 0 {
		fmt.Println("\nNão é possível realizar um depósito negativo...\nVerifique e tente novamente.\n__________________________________")
		time.Sleep(3 * time.Second)
	}

}
func (c *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) {
	transferencia := valor > 0 && valor <= c.saldo
	if transferencia {
		c.saldo -= valor
		contaDestino.Depositar(valor)
		fmt.Println("\nTransferencia realizada no valor de R$:", valor, "\nSeu saldo atual:", c.saldo)
		fmt.Print("saldo do desitnatario R$:", contaDestino.saldo)
		time.Sleep(3 * time.Second)
	} else {
		fmt.Println("\nOperação inválida...\nVerifique e tente novamente.")
		time.Sleep(3 * time.Second)
	}
}
func (c *ContaCorrente) Exibirsaldo() {
	fmt.Println("\n\nsaldo R$:", c.saldo, "\n__________________________________")
	time.Sleep(3 * time.Second)
}
func (c *ContaCorrente) CadNovaConta() {
	jsonFile, err := os.Open(`contas/contas.json`)

	if err != nil {
		//Caso tenha tido erro, ele é apresentado na tela
		fmt.Println(err)
	}
	fmt.Println("Digite o nome do Titular")
	fmt.Scan(&c.Titular.Nome)
	fmt.Println("O nome do titular é:", c.Titular.Nome)
	fmt.Println("Digite o CPF do titular")
	fmt.Scan(&c.Titular.Cpf)
	fmt.Println("Digite a profissão do titular")
	fmt.Scan(&c.Titular.Profissao)

	for {
		fmt.Println("Digite uma senha númerica:")
		var pass int
		fmt.Scan(&pass)
		var confpass int
		fmt.Println("Congirme a senha:")
		fmt.Scan(&confpass)
		if pass == confpass {
			c.Titular.CadSenha(pass)
			fmt.Println("Senha cadastrada com sucesso")
			fmt.Println(c.Titular.VerificarSenha())

			break
		} else {
			fmt.Println("As senhas não conferem, tente novamente.")
		}
	}

	c.NumAgencia = 354
	c.Operacao = 2
	c.NumConta = int(rand.Int31()) + rand.Intn(1200) + rand.Intn(1000000)
	fmt.Println("conta gerada:", c)
	//////nome, pass := c.Titular.RegistrarConta()
	//Aqui vamos converter nosso objBook com o nome alterado em bytes
	//////c.Titular.Nome = nome
	byteValueJSON, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
	}

	//Por fim vamos salvar em um arquivo JSON alterado.
	err = ioutil.WriteFile("contas/c.json", byteValueJSON, 0644)
	if err != nil {
		fmt.Println(err)
	}
	jsonFile.Close()
}
