package models

type Art struct {
}

func (a *Art) CriarCadeira() ICadeira {
	return &ArtCadeira{
		Cadeira: Cadeira{
			Descricao: "Cadeira",
			Style:     "Art",
		},
	}
}
func (a *Art) CriarMesa() IMesa {
	return &ArtMesa{
		Mesa: Mesa{
			Descricao: "Mesa",
			Style:     "Art",
		},
	}
}
func (a *Art) CriarPoltrona() IPoltrona {
	return &ArtPoltrona{
		Poltrona: Poltrona{
			Descricao: "Poltrona",
			Style:     "Art",
		},
	}
}
func (a *Art) CriarArmario() IArmario {
	return &ArtArmario{
		Armario: Armario{
			Descricao: "Armario",
			Style:     "Art",
		},
	}
}
func (a *Art) CriarCama() ICama {
	return &ArtCama{
		Cama: Cama{
			Descricao: "Cama",
			Style:     "Art",
		},
	}
}
