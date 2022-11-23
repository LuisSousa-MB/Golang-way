package exercicios

import (
	"fmt"
	"os"
)

var notas map[string][]int = make(map[string][]int)

func addToMap(nome string, nota int) {
	notas[nome] = append(notas[nome], nota)
	fmt.Println(&notas)

}

func main() {
	fmt.Println("endereço de notas", &notas, "notas:", notas)
	ExibirMenu()
}

func ExibirMenu() {
	for i := 0; i < 1; {

		fmt.Println("Bem vindo!\n\nSelecione a opção:\n\n1 - Adicionar Aluno e Notas\n2 - Verificar notas e média de Aluno\n3 - Todos os Alunos e médias\n0 - Sair")
		var op int
		fmt.Scan(&op)
		SelecionarOpcao(op)

	}

}
func SelecionarOpcao(op int) {
	var nome string
	var nota int
	media := 0.0
	mediaG := 0.0
	switch op {
	case 1:
		fmt.Println("informe o nome do aluno:")
		fmt.Scan(&nome)
		for i := 1; i <= 4; i++ {
			fmt.Printf("Informe a %dª nota:", i)
			fmt.Scan(&nota)
			addToMap(nome, nota)
		}
	case 2:
		fmt.Println("Buscar média por aluno\n\nInsira o nome do aluno para a busca:")
		fmt.Scan(&nome)
		fmt.Println("\nEstas são as notas e a média do aluno", nome)
		for i := 0; i <= 3; i++ {
			media += float64(notas[nome][i])
			fmt.Printf("%dª nota: %d\n", i+1, notas[nome][i])
		}
		media = media / 4
		fmt.Printf("Média: %2.f\n\n", media)
	case 3:
		fmt.Println("Imprimindo todos os Alunos:\n\n")
		for i := 1; i < len(notas); i++ {
			for a := range notas {
				fmt.Printf("Aluno %s\n\n", a)
				for i := 0; i <= 3; i++ {
					media += float64(notas[a][i])
					fmt.Printf("%dª nota: %d\n", i+1, notas[a][i])
				}
				media = media / 4
				mediaG += media
				fmt.Printf("Média: %2.f\n\n", media)

			}
			mediaG = mediaG / float64(len(notas))
			fmt.Printf("\nMédia geral dos alunos: %2.f\n\n", mediaG)

		}
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0) // Envia o status "0", que significa sair com sucesso, já o stautus "-1" indicaria um possivel problema/erro.
	default:
		fmt.Println("Comando inválido..")

	}
}
