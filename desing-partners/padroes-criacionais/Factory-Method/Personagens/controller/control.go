package controller

import (
	"Personagens-RPG/models"
	"fmt"
	rand "math/rand"
	"time"
)

func GeradorDePersonagem(name string) *models.IPersonagem {
	fmt.Println("Gerando novo personagem...")
	novoPersongem := defineClasse(randomizadorPersonagem())
	return novoPersongem
}
func ExibePersonagem(p models.IPersonagem) {
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
	fmt.Printf("Agilidade: %d", p.GetAgilidade())
	fmt.Println()
	fmt.Printf("Level: %d", p.GetLevel())
	fmt.Println()
	fmt.Printf("XP: %d", p.GetXP())
	fmt.Println()
}
func randomizadorPersonagem() (name string, hp, mp, poder, forca, agilidade, level, xp int) {
	rand.Seed(time.Now().UnixNano())
	names := [30]string{"Gobu", "Taos", "Bekye", "Cyeno", "Geada", "Ankey", "Wuema", "Rofes", "Pezea",
		"Tetel", "Cles", "Tifaytal", "Zeage", "Keate", "Lupia", "Ulel", "Curumu", "Rulas", "Caferius", "Sebasth", "Lorpheus", "Mizandre", "Saujin",
		"Luna", "Partenes", "Teféia", "Zanza", "Penlia", "Aziane", "Garrata"}
	nameRand := rand.Intn(len(names))
	name = names[nameRand]
	hp = 200 + rand.Intn(200)
	mp = 0
	poder = 10
	forca = 10
	agilidade = 10
	level = 1
	xp = 0
	points := 21
	for points > 0 {
		fmt.Println("pontos atuais", points)
		dado := 1 + rand.Intn(3)
		fmt.Println("Dados de pontos:", dado)
		time.Sleep(500 * time.Millisecond)
		dado2 := 1 + rand.Intn(3)
		switch dado2 {
		case 1:
			if dado <= points {
				fmt.Println("Points em Poder:", dado)
				time.Sleep(500 * time.Millisecond)
				poder += dado
				points -= dado
			}
		case 2:
			if dado <= points {
				fmt.Println("Points em Força:", dado)
				time.Sleep(500 * time.Millisecond)
				forca += dado
				points -= dado
			}
		case 3:
			if dado <= points {
				fmt.Println("Points em Agilidade:", dado)
				time.Sleep(500 * time.Millisecond)
				agilidade += dado
				points -= dado
			}
		}
	}
	mp = poder * 12
	return name, hp, mp, poder, forca, agilidade, level, xp
}

func defineClasse(name string, hp, mp, poder, forca, agilidade, level, xp int) *models.IPersonagem {
	rand.Seed(time.Now().UnixNano())
	equilibriumx := 2
	if forca > agilidade && forca > poder {
		forca, poder = equilibrium(forca, poder)
		novoGuerreiro := models.NewGuerreiro(name, hp, mp, poder, forca, agilidade, level, xp)
		return &novoGuerreiro

	} else if agilidade > poder && agilidade > forca {
		agilidade, poder = equilibrium(agilidade, poder)
		novoArqueiro := models.NewArqueiro(name, hp, mp, poder, forca, agilidade, level, xp)
		return &novoArqueiro
	} else if poder > agilidade && poder > forca {
		poder, forca = equilibrium(poder, forca)
		novoMago := models.NewMago(name, hp, mp, poder, forca, agilidade, level, xp)
		return &novoMago
	} else {
		luck := 1 + rand.Intn(3)
		switch luck {
		case 1:
			forca, poder = equilibrium(forca+equilibriumx, poder-equilibriumx)
			novoGuerreiro := models.NewGuerreiro(name, hp, mp, poder, forca, agilidade, level, xp)
			return &novoGuerreiro
		case 2:
			agilidade, poder = equilibrium(agilidade+equilibriumx, poder-equilibriumx)
			novoArqueiro := models.NewArqueiro(name, hp, mp, poder, forca, agilidade, level, xp)
			return &novoArqueiro
		case 3:
			poder, forca = equilibrium(poder+equilibriumx, forca-equilibriumx)
			novoMago := models.NewMago(name, hp, mp, poder, forca, agilidade, level, xp)
			return &novoMago
		}
		novoGuerreiro := models.NewGuerreiro(name, hp, mp, poder, forca, agilidade, level, xp)
		return &novoGuerreiro
	}
}

func equilibrium(alfa, beta int) (maior, menor int) {
	rand.Seed(time.Now().UnixNano())
	equilibrium := rand.Intn(5) + 2
	fmt.Println("Equilibrando distribuição em:", equilibrium)
	beta -= equilibrium
	alfa += equilibrium
	return alfa, beta
}
