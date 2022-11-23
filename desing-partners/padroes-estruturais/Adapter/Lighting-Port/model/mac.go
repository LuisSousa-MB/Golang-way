package model

import "fmt"

type Mac struct {
}

func (m *Mac) InserirNoConectorLighting() {
	fmt.Println("O conector Lighting foi acoplado no computador MAC.")
}
