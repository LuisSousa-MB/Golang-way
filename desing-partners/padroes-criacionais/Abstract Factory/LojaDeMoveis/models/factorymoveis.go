package models

import (
	"fmt"
)

type IMoveisFacotory interface {
	CriarCadeira() ICadeira
	CriarMesa() IMesa
	CriarPoltrona() IPoltrona
	CriarArmario() IArmario
	CriarCama() ICama
}

func GetMaterialStyle(style string) (IMoveisFacotory, error) {
	if style == "Art" || style == "art" {
		return &Art{}, nil
	} else if style == "Vitoriano" || style == "vitoriano" {
		return &Vitoriano{}, nil
	} else if style == "Moderno" || style == "moderno" {
		return &Moderno{}, nil
	} else {
		return nil, fmt.Errorf("Estilo não encontrado...")
	}
}
