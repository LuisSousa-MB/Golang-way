package models

type Moderno struct {
}

func (a *Moderno) CriarCadeira() ICadeira {
	return &ModernoCadeira{
		Cadeira: Cadeira{
			Descricao: "Cadeira",
			Style:     "Moderno",
		},
	}
}
func (a *Moderno) CriarMesa() IMesa {
	return &ModernoMesa{
		Mesa: Mesa{
			Descricao: "Mesa",
			Style:     "Moderno",
		},
	}
}
func (a *Moderno) CriarPoltrona() IPoltrona {
	return &ModernoPoltrona{
		Poltrona: Poltrona{
			Descricao: "Poltrona",
			Style:     "Moderno",
		},
	}
}
func (a *Moderno) CriarArmario() IArmario {
	return &ModernoArmario{
		Armario: Armario{
			Descricao: "Armamario",
			Style:     "Moderno",
		},
	}
}
func (a *Moderno) CriarCama() ICama {
	return &ModernoCama{
		Cama: Cama{
			Descricao: "Cama",
			Style:     "Moderno",
		},
	}
}
