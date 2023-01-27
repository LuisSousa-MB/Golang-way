package models

import (
	"math/rand"
	"time"
)

type IHabilidades interface {
}
type Effect interface {
}

type Habilidade struct {
	Descrição string
	Nome      string
	Dano      int
	Efeito    Effect
}

func (p *Personagem) BolaDeFogo(level, poder int) int {
	rand.Seed(time.Now().UnixNano())
	dano := 0
	if level < 3 {
		p.MP -= 40
		dano = poder*(rand.Intn((poder*(rand.Intn(3))))-poder) + poder
	} else if level >= 3 && level <= 6 {
		p.MP -= 55
		dano = poder*(rand.Intn((poder*(rand.Intn(5))))-poder) + poder
	} else if level >= 7 {
		p.MP -= 69
		dano = poder*(rand.Intn((poder*(rand.Intn(9))))-poder) + poder
	}

	return dano
}
