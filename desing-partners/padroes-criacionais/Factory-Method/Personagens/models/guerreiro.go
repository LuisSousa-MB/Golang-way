package models

type Guerreiro struct {
	Personagem
}

func NewGuerreiro(name string, hp, mp, poder, forca, defesa, agilidade, level, xp int) IPersonagem {
	return &Guerreiro{
		Personagem: Personagem{
			Class:     "Guerreiro",
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
