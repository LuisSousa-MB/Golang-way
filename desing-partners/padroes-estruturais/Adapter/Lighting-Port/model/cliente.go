package model

import "fmt"

type Cliente struct {
}

func (c *Cliente) InserirNoConectorLightingDoComputador(pc Computador) {
	fmt.Println("O cliente inseriu o conector Lighting no computador...")
	pc.InserirNoConectorLighting()
}
