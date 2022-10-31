package models

type IConstrutor interface {
	SetQuantidadeJanelas()
	SetEstiloDeJanelas()
	SetEstiloDoPiso()
	SetEstiloDePorta()
	SetQuantidadeDeAndares()
	SetChamine()
	GetConstrucao() Construcao
}

func GetConstrutor(tipoDeConstrutor string) IConstrutor {
	if tipoDeConstrutor == "Normal" {
		return novoConstrutorNormal()
	}
	if tipoDeConstrutor == "Moderno" {
		return novoConstrutorModerno()
	}
	return nil
}
