package controller

import (
	"Personagens-RPG/database"
	"Personagens-RPG/models"
	"fmt"
	rand "math/rand"
	"time"
)

func InsertPersonagemDB(personagem models.IPersonagem) {
	newPersonagem := models.Personagem{Class: personagem.GetClass(), Name: personagem.GetName(),
		HP: personagem.GetHP(), MP: personagem.GetMP(), Poder: personagem.GetPoder(), Forca: personagem.GetForca(),
		Defesa: personagem.GetDefesa(), Agilidade: personagem.GetAgilidade(), Level: personagem.GetLevel(), XP: personagem.GetXP(),
	}
	/* 	result := database.DB.Last(&newPersonagem)
	   	fmt.Println("Ultimo ID personagem", result)
	   	fmt.Println("Primeiro personagem gerado, ID:", personagem.GetID())

	   	fmt.Println(newPersonagem.GetID(), "ID capturado")
	   	fmt.Println("id error", "comment") */
	database.DB.Create(&newPersonagem)

}

func GeradorDePersonagem() *models.IPersonagem { //Publica. Gerencia a ordem das funções, cria e retorna o personagem.
	fmt.Println("Gerando novo personagem...")
	novoPersongem := defineClasse(randomizadorPersonagem())
	InsertPersonagemDB(*novoPersongem)
	return novoPersongem
}
func randomizadorPersonagem() (name string, hp, mp, poder, forca, defesa, agilidade, level, xp int) { // Cria de maneira randomica o nome e os atributos para o personagem
	rand.Seed(time.Now().UnixNano())
	names := []string{"Gobu", "Taos", "Bekye", "Cyeno", "Geada", "Ankey", "Wuema", "Rofes", "Pezea",
		"Tetel", "Cles", "Tifaytal", "Zeage", "Keate", "Lupia", "Ulel", "Curumu", "Rulas", "Caferius", "Sebasth", "Lorpheus", "Mizandre", "Saujin",
		"Luna", "Partenes", "Teféia", "Zanza", "Penlia", "Aziane", "Garrata"}
	//fmt.Println("tamanho array: ", len(names))
	nameRand := rand.Intn(len(names) - 1)
	name = names[nameRand]
	hp = 200 + rand.Intn(200)
	level = 1
	poder = 10
	forca = 10
	agilidade = 10
	xp = 0
	points := 21
	for points > 0 {
		fmt.Println("pontos atuais", points)
		dado := 1 + rand.Intn(3)
		fmt.Println("Dados de pontos:", dado)
		time.Sleep(250 * time.Millisecond)
		dado2 := 1 + rand.Intn(3)
		switch dado2 {
		case 1:
			if dado <= points {
				fmt.Println("Points em Poder:", dado)
				time.Sleep(250 * time.Millisecond)
				poder += dado
				points -= dado
			}
		case 2:
			if dado <= points {
				fmt.Println("Points em Força:", dado)
				time.Sleep(250 * time.Millisecond)
				forca += dado
				points -= dado
			}
		case 3:
			if dado <= points {
				fmt.Println("Points em Agilidade:", dado)
				time.Sleep(250 * time.Millisecond)
				agilidade += dado
				points -= dado
			}
		}
	}

	return name, hp, mp, poder, forca, defesa, agilidade, level, xp
}

func defineClasse(name string, hp, mp, poder, forca, defesa, agilidade, level, xp int) *models.IPersonagem { //Retorna persoagem completo e define sua classe de acordo com a distribuição de pontos
	rand.Seed(time.Now().UnixNano())
	equilibriumx := 2
	if forca > agilidade && forca > poder {
		forca, poder = equilibrium(forca, poder)
		mp = poder * 12
		defesa = ((forca * 3) + (agilidade / 2)) / 4
		novoGuerreiro := models.NewGuerreiro(name, hp, mp, poder, forca, defesa, agilidade, level, xp)
		return &novoGuerreiro

	} else if agilidade > poder && agilidade > forca {
		agilidade, poder = equilibrium(agilidade, poder)
		mp = poder * 12
		defesa = ((forca * 3) + (agilidade / 2)) / 4
		novoArqueiro := models.NewArqueiro(name, hp, mp, poder, forca, defesa, agilidade, level, xp)
		return &novoArqueiro
	} else if poder > agilidade && poder > forca {
		poder, forca = equilibrium(poder, forca)
		mp = poder * 12
		defesa = ((forca * 3) + (agilidade / 2)) / 4
		novoMago := models.NewMago(name, hp, mp, poder, forca, defesa, agilidade, level, xp)
		return &novoMago
	} else {
		luck := 1 + rand.Intn(3)
		switch luck {
		case 1:
			forca, poder = equilibrium(forca+equilibriumx, poder-equilibriumx)
			mp = poder * 12
			defesa = ((forca * 3) + (agilidade / 2)) / 4
			novoGuerreiro := models.NewGuerreiro(name, hp, mp, poder, forca, defesa, agilidade, level, xp)
			return &novoGuerreiro
		case 2:
			agilidade, poder = equilibrium(agilidade+equilibriumx, poder-equilibriumx)
			mp = poder * 12
			defesa = ((forca * 3) + (agilidade / 2)) / 4
			novoArqueiro := models.NewArqueiro(name, hp, mp, poder, forca, defesa, agilidade, level, xp)
			return &novoArqueiro
		case 3:
			poder, forca = equilibrium(poder+equilibriumx, forca-equilibriumx)
			mp = poder * 12
			defesa = ((forca * 3) + (agilidade / 2)) / 4
			novoMago := models.NewMago(name, hp, mp, poder, forca, defesa, agilidade, level, xp)
			return &novoMago
		}
		mp = poder * 12
		defesa = ((forca * 3) + (agilidade / 2)) / 4
		novoGuerreiro := models.NewGuerreiro(name, hp, mp, poder, forca, defesa, agilidade, level, xp)
		return &novoGuerreiro
	}
}

func equilibrium(alfa, beta int) (maior, menor int) { //Corrige possivel desequilibrio da distribuição de pontos conforme a classe
	rand.Seed(time.Now().UnixNano())
	equilibrium := rand.Intn(5) + 2
	fmt.Println("Equilibrando distribuição em:", equilibrium)
	beta -= equilibrium
	alfa += equilibrium
	return alfa, beta
}
func ExibePersonagem(p models.IPersonagem) { // Exibe os dados do personagem criado
	fmt.Println()
	fmt.Println("Personagem gerado!!!")
	fmt.Println()
	fmt.Printf("Nome: %s", p.GetName())
	fmt.Println()
	fmt.Printf("Classe: %s", p.GetClass())
	fmt.Println()
	fmt.Printf("HP: %d", p.GetHP())
	fmt.Println()
	fmt.Printf("MP: %d", p.GetMP())
	fmt.Println()
	fmt.Printf("Poder: %d", p.GetPoder())
	fmt.Println()
	fmt.Printf("Força: %d", p.GetForca())
	fmt.Println()
	fmt.Printf("Defesa: %d", p.GetDefesa())
	fmt.Println()
	fmt.Printf("Agilidade: %d", p.GetAgilidade())
	fmt.Println()
	fmt.Printf("Level: %d", p.GetLevel())
	fmt.Println()
	fmt.Printf("XP: %d", p.GetXP())
	fmt.Println()
}
