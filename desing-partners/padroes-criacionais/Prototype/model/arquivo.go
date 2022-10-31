package model

import (
	"Prototype/prototype"
	"fmt"
)

type Arquivo struct {
	Nome string
}

func (a *Arquivo) Exibir(indentation string) {
	fmt.Println(indentation + a.Nome)
}

func (a *Arquivo) Clonar() prototype.Inode {
	return &Arquivo{Nome: a.Nome + "_copia"}
}
