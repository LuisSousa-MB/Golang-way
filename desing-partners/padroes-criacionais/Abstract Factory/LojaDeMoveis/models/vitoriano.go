package models

type Vitoriano struct {
}

func (a *Vitoriano) CriarCadeira() ICadeira {
	return &VitorianoCadeira{
		Cadeira: Cadeira{
			Descricao: "Cadeira",
			Style:     "Vitoriano",
		},
	}
}
func (a *Vitoriano) CriarMesa() IMesa {
	return &VitorianoMesa{
		Mesa: Mesa{
			Descricao: "Mesa",
			Style:     "Vitoriano",
		},
	}
}
func (a *Vitoriano) CriarPoltrona() IPoltrona {
	return &VitorianoPoltrona{
		Poltrona: Poltrona{
			Descricao: "Poltrona",
			Style:     "Vitoriano",
		},
	}
}
func (a *Vitoriano) CriarArmario() IArmario {
	return &VitorianoArmario{
		Armario: Armario{
			Descricao: "Armario",
			Style:     "Vitoriano",
		},
	}
}
func (a *Vitoriano) CriarCama() ICama {
	return &VitorianoCama{
		Cama: Cama{
			Descricao: "Cama",
			Style:     "Vitoriano",
		},
	}
}
