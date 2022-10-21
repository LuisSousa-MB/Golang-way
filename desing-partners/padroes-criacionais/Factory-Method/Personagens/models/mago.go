package models

type Mago struct {
	Personagem
}

func NewMago(name string, hp, mp, poder, forca, agilidade, level, xp int) IPersonagem {
	return &Guerreiro{
		Personagem: Personagem{
			Class:     "Mago",
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
