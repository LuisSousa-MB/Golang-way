package models

type Mago struct {
	Personagem
}

func NewMago(name string, hp, mp, poder, forca, defesa, agilidade, level, xp int) IPersonagem {
	return &Mago{
		Personagem: Personagem{
			Class:     "Mago",
			Name:      name,
			HP:        hp,
			MP:        mp,
			Defesa:    defesa,
			Poder:     poder,
			Forca:     forca,
			Agilidade: agilidade,
			Level:     level,
			XP:        xp,
		},
	}
}
