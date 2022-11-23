package codes

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"time"
)

var NewChain Chain
var Tx Transaction
var TxPoll Data
var NewBlock Block
var dificuldade int = 1

func FormatHash(h string) string {
	if len(h) > 64 {
		cut := len(h) - 64
		h = h[:len(h)-cut]
		//println("new len: ", len(h))
		//println("Formatedhash: ", h)
	}
	if len(h) < 64 {
		need := 64 - len(h)
		for i := 0; i < need; i++ {
			x := rand.Intn(122-48+1) + 48
			x = RangeOfHash(x)
			h += string(x)

		}
		//println("new len: ", len(h))
		//println("Formatedhash: ", h)
	}
	return h
}

func RangeOfHash(x int) int {
	if x >= 58 && x <= 64 {
		x += 7
		return x
	}
	if x >= 91 && x <= 96 {
		x += 6
		return x
	}
	return x
}

func GenerateHashByte(pass string) string {
	rand.Seed(777)

	hash := ""
	cofre := make([]int, 5)
	//fmt.Println("pass string to hash:", pass)
	for _, h := range pass {
		//println("h:", h)

		cofre = append(cofre, int(h))
	}

	//fmt.Printf("confre :%v\n", cofre)
	secretGenerator := func() int {
		var segredo int

		for i := 0; i < len(cofre); i++ {

			//fmt.Printf("Cofre na posição %d: %d\n", i, cofre[i])
			multiSecret := rand.Intn(77777777)
			segredo = segredo + cofre[i]*multiSecret
		}
		//fmt.Println(segredo)
		return segredo
	}

	for i := 0; i < len(cofre); i++ {
		rand.Seed(int64(secretGenerator()))
		//println("Secret Key:", secretGenerator())
		x := rand.Intn(122-48+1) + 48
		x = RangeOfHash(x)
		//println("X: ", x)
		hash += string(rune(x))
	}

	//println("Firsthash: ", hash)
	//println("len: ", len(hash))

	return FormatHash(hash)
}

func ExibeMenu() {
	for i := 0; i == 0; {
		//fmt.Printf("Tamanho da chain é %d\n", len(NewChain.Instance))

		var op int
		fmt.Println("Bem vindo a sua Go BlockChain!")
		fmt.Println()
		fmt.Println("",
			"1 - Minerar Blockchain\n",
			"2 - Exibir Blockchain\n",
			"3 - Validar Blockchain\n",
			"4 - Adicinoar Transação\n",
			"5 - Exibir Transações do bloco atual\n",
			"6 - Exibir dados pro bloco\n",
			"7 - Alterar dificuldade da rede\n",
			"8 - Configurar mineradores\n",
			"9 - Testar Hash\n",
			"")
		fmt.Scan(&op)
		SelecionarOpcao(op)
	}
}
func SelecionarOpcao(op int) {
	switch op {
	case 1:
		NewChain.GerarBlock(dificuldade)
	case 2:
		PrintChain()
	case 3:
		IsValid()
	case 4:
		RegisterTransaction()
	case 5:
		ShowOpenTransactions()
	case 6:
		ShowBlockInfo()
	case 7:
		SetDificult()
	case 8:
		ConfigMiners()

	case 9:
		ProvarHash()
	}
}
func (c *Chain) PrintData(i int) {

	for j := 0; j < len(c.Instance[i].Data.Transactions); j++ {
		fmt.Printf("Transação : %d", j+1)
		fmt.Println()
		fmt.Println("\tRemetente: ", c.Instance[i].Data.Transactions[j].Sender)
		fmt.Println("\tDestinatario: ", c.Instance[i].Data.Transactions[j].Receiver)
		fmt.Println("\tNota: ", c.Instance[i].Data.Transactions[j].Msg)
		fmt.Println("\tValor: ", c.Instance[i].Data.Transactions[j].Value)
		fmt.Println("\tData: ", c.Instance[i].Data.Transactions[j].TimeStamp)
		fmt.Println("\tHash: ", c.Instance[i].Data.Transactions[j].Hash)

	}

}
func (c *Chain) PrintBlock(i int) {
	var sTimestamp string
	fmt.Println("Index:", NewChain.Instance[i].Index)
	fmt.Println("Nonce:", NewChain.Instance[i].Nonce)
	fmt.Println("Hash:", NewChain.Instance[i].Hash)
	fmt.Println("Previous Hash:", NewChain.Instance[i].PreviousHash)
	sTimestamp = NewChain.Instance[i].TimeStamp.Format("2006-01-02 15:04:05")
	fmt.Println("TimeStamp:", sTimestamp)
	//fmt.Println("Data:", NewChain.Instance[i].Data)
	fmt.Println("Data:")
	NewChain.PrintData(i)
	fmt.Println("_______________________________________________")

}
func IsValid() {
	if len(NewChain.Instance) == 0 {
		fmt.Println("Blockchain has not been initialized")
		time.Sleep(time.Second * 3)

	} else {
		valid := false
		var hashs []string
		for i := 0; i < len(NewChain.Instance); i++ {
			for pos, e := range hashs {
				for j := pos + 1; j < len(hashs); j++ {

					if e == hashs[j] {
						fmt.Println("Hash violado! Repetido!!!")
						valid = false
						time.Sleep(time.Second * 3)

					}
				}
				hashs = append(hashs, NewChain.Instance[i].Hash)
			}
			tamanho := len(NewChain.Instance)
			fmt.Println()
			for i := 0; i < tamanho-1; i++ {
				if NewChain.Instance[i].Hash == NewChain.Instance[i+1].PreviousHash {
					//fmt.Printf("Block %d is Ok!\n", i+2)
					valid = true
				} else {
					valid = false
				}
			}

		}
		if valid {
			fmt.Println("BlockChain is Integral!!!")
			time.Sleep(time.Second * 3)

		} else {
			fmt.Println("BlockChain is Violated!!!")
			time.Sleep(time.Second * 3)

		}
		time.Sleep(time.Second * 2)
		fmt.Println()
	}
}
func RegisterTransaction() {
	var msg, sender, receiver, hash string
	var value float64
	var timeStamp time.Time
	fmt.Print("Informe o nome do remetente: ")
	fmt.Scan(&sender)
	fmt.Print("Informe o nome do destinatario: ")
	fmt.Scan(&receiver)
	fmt.Print("Escreva uma nota: ")
	fmt.Scan(&msg)
	fmt.Print("Informe o valor da transação: ")
	fmt.Scan(&value)
	timeStamp = time.Now()
	hash = GenerateHashByte(timeStamp.String() + sender + receiver + msg + strconv.FormatFloat(value, 'E', -1, 64))
	Tx = Transaction{
		Msg:       msg,
		TimeStamp: timeStamp.Format("2006-01-02 15:04:05"),
		Sender:    sender,
		Receiver:  receiver,
		Hash:      hash,
		Value:     value,
	}
	TxPoll.Transactions = append(TxPoll.Transactions, Tx)
}
func ShowOpenTransactions() {
	for j := 0; j < len(TxPoll.Transactions); j++ {
		fmt.Println()
		fmt.Println("\tRemetente: ", TxPoll.Transactions[j].Sender)
		fmt.Println("\tDestinatario: ", TxPoll.Transactions[j].Receiver)
		fmt.Println("\tNota: ", TxPoll.Transactions[j].Msg)
		fmt.Println("\tValor: ", TxPoll.Transactions[j].Value)
		fmt.Println("\tData: ", TxPoll.Transactions[j].TimeStamp)
		fmt.Println("\tHash: ", TxPoll.Transactions[j].Hash)
		fmt.Println("_______________________________________________________")
	}
}

func SetDificult() {
	var setDificuldade uint
	fmt.Printf("A dificuldade atual está em %d.\nDigite a nova dificuldade: \n", dificuldade)
	fmt.Scan(&setDificuldade)
	if setDificuldade <= 32 {
		dificuldade = int(setDificuldade)
		fmt.Println("Dificuldade alterada com sucesso.\nA nova dificuldade é ", dificuldade)
	} else {
		fmt.Println("Dificuldade inválida!!!")
	}
}
func PrintChain() {
	tamanho := len(NewChain.Instance)
	fmt.Println(NewChain.Instance)
	fmt.Println()
	for i := 0; i < tamanho; i++ {
		NewChain.PrintBlock(i)
	}
	time.Sleep(time.Second * 3)
}

func ConfigMiners() {
	fmt.Println("O número máximo de mineradores recomendados para o seu sistema é ", runtime.NumCPU())
	fmt.Println("Você está utilizando ", Miners, " Mineradores.")
	newMiners := 0
	for i := 0; i < 1; {
		fmt.Println("Quantos mineradores deseja utilizar?")
		fmt.Scan(&newMiners)
		if newMiners > 0 && newMiners <= 8 {
			i++
			fmt.Println("Novo valor configurado!")
			Miners = newMiners
			fmt.Println("Você está utilizando ", Miners, " Mineradores.")
			time.Sleep(time.Second * 3)

		} else {
			fmt.Println("Valor inválido ou superior ao permitido pelo seu sistema atual, tente um valor entre 1 e", runtime.NumCPU())
		}
	}
}
func ProvarHash() {
	/* pass := ""
	fmt.Println("Informe uma cadeia de caracteres para transformar em GoHash: ")
	fmt.Scan(&pass)
	hashpass1 := GenerateHashByte(pass)
	fmt.Println("Hash gerado: ", hashpass1)
	time.Sleep(time.Second * 3)
	fmt.Println("Gerando novamente... ")
	time.Sleep(time.Second * 2)
	hashpass2 := GenerateHashByte(pass)
	fmt.Println("Hash gerado: ", hashpass2)
	time.Sleep(time.Second * 2)
	fmt.Println()
	if hashpass1 == hashpass2 {
		fmt.Println("Sucesso!")
	} else {
		fmt.Println("Falha...")
	}
	time.Sleep(time.Second * 2)
	fmt.Println() */

	chaveTest := "047160000000000GeNeSiS000000000000000Wed Nov 23 20:14:10 -03 2022"
	hashTest := "GowG9oqUSvthfl9U5YFlFuFfdCwhVfgbicaPACsa1UbPDbPf2F4O8fXtGI1ldPqB"

	fmt.Println("Testando a chave salva no sistema...\n", chaveTest, "\nO resultado deve ser:", hashTest)
	time.Sleep(time.Second * 3)
	fmt.Println()

	hash := GenerateHashByte(chaveTest)
	fmt.Println(hash)
	time.Sleep(time.Second * 1)
	fmt.Println()

	if hash == hashTest {
		fmt.Println("Sucesso!")
	} else {
		fmt.Println("Falha...")
	}
	time.Sleep(time.Second * 3)
}
