package models

type Arqueiro struct {
	Personagem
}

func NewArqueiro(name string, hp, mp, poder, forca, defesa, agilidade, level, xp int) IPersonagem {
	return &Arqueiro{
		Personagem: Personagem{
			Class:     "Arqueiro",
			Name:      name,
			HP:        hp,
			MP:        mp,
			Poder:     poder,
			Defesa:    defesa,
			Forca:     forca,
			Agilidade: agilidade,
			Level:     level,
			XP:        xp,
		},
	}
}
