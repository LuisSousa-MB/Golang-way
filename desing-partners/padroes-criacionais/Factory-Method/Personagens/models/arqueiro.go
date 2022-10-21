package models

type Arqueiro struct {
	Personagem
}

func NewArqueiro(name string, hp, mp, poder, forca, agilidade, level, xp int) IPersonagem {
	return &Guerreiro{
		Personagem: Personagem{
			Class:     "Arqueiro",
			Name:      name,
			HP:        hp,
			MP:        mp,
			Poder:     poder,
			Forca:     forca,
			Agilidade: agilidade,
			Level:     level,
			XP:        xp,
		},
	}
}
