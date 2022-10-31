package model

import (
	"Prototype/prototype"
	"fmt"
)

type Pasta struct {
	Filho []prototype.Inode
	Nome  string
}

func (p *Pasta) Exibir(indentation string) {
	fmt.Println(indentation + p.Nome)
	for _, i := range p.Filho {
		i.Exibir(indentation + indentation)
	}
}

func (p *Pasta) Clonar() prototype.Inode {
	ClonarPasta := &Pasta{Nome: p.Nome + "_copia"}
	var tempFilho []prototype.Inode
	for _, i := range p.Filho {
		copia := i.Clonar()
		tempFilho = append(tempFilho, copia)
	}
	ClonarPasta.Filho = tempFilho
	return ClonarPasta
}
