package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 2
const intervaloMinutos = 1

func main() {
	exibeIntro()
	for {
		exibeMenu()
		comando := lerComando()

		/*if comando == 1 {
			fmt.Println("Monitorando...")
		} else if comando == 2 {
			fmt.Println("Exibindo Logs:")
		} else if comando == 0 {
			fmt.Println("Saindo do programa...")
		} else {
			fmt.Println("Comando inválido..")
		}*/
		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs:")
			println("")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0) // Envia o status "0", que significa sair com sucesso, já o stautus "-1" indicaria um possivel problema/erro.
		default:
			fmt.Println("Comando inválido..")
		}
	}

}

func exibeIntro() {
	nome := "Henrique"
	var versao float32 = 1.1
	fmt.Println("Olá, Sr.(a) ", nome)
	fmt.Println("Este programa está na versão ", versao)
	println("")
}

func exibeMenu() {
	fmt.Println(("1- Iniciar monitoramento "))
	fmt.Println(("2- Exibir Logs"))
	fmt.Println(("0 - Sair do programa... "))
}

func lerComando() int {

	var comando int
	//fmt.Scanf("%d", &comando) %d = decimal (aceitar inteiros)
	fmt.Scan(&comando)
	//fmt.Println("O comando selecionado foi", comando)

	return comando
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	println("")

	/* var sites [3] string

	sites[0] = "https://www.linkedin.com/"
	sites[1] = "https://www.facebook.com/"
	sites[2] = "https://www.twitter.com/" */ //metódo com array
	//sites := []string{"https://www.linkedin.com/", "https://www.facebook.com/", "https://www.twitter.com/"} // Agora usamos o slice do arquivo sites.txt

	sites := leArquivoDeSites()
	for i := 0; i < monitoramentos; {
		println("Monitoramento", i+1, "de", monitoramentos)
		println("______________________________________________")
		i++
		for _, site := range sites {
			testaSite(site)
			println("______________________________________________")

		}
		if i < monitoramentos {
			println("Aguardando para o próximo monitoramento...")
			time.Sleep(intervaloMinutos * time.Minute)
			println("______________________________________________")
		}

	}

}

func testaSite(site string) {
	resp, err := http.Get(site)
	println("Testando o site:", site)

	if err != nil {
		fmt.Println("Erro ao receber resposta do http.Get:", err)
	}

	if resp.StatusCode == 200 {
		println("Status: 200 | Carregado com sucesso!")
		registraLog(site, true, resp.StatusCode)
	} else {
		println("Status: ", resp.StatusCode, "| Erro ao carregar o site:", site)
		registraLog(site, false, resp.StatusCode)
	}
	time.Sleep(1 * time.Second)
}

func leArquivoDeSites() []string {

	var sites []string

	sitestxt, err := os.Open("sites.txt")

	//sitestxt, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		println("Erro ao ler arquivo | Erro:", err)
	}

	leitor := bufio.NewReader(sitestxt)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		//fmt.Println(linha)

		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}

	//fmt.Println(sites)
	sitestxt.Close()
	return sites
}

func registraLog(site string, status bool, code int) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Erro ao registrar logs:", err)
	}
	stringcode := strconv.Itoa(code)

	arquivo.WriteString(site + " - online:" + strconv.FormatBool(status) + " StatusCode: " + stringcode + "\n" + "Data: " + time.Now().Format("02/01/2006 15:04:05") + "\n" +
		"________________________________________________________" + "\n")

	arquivo.Close()

}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Erro ao exibir logs: ", err)
	}

	fmt.Println(string(arquivo))

}
