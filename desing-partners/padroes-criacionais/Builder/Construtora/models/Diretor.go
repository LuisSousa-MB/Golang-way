package models

type Diretor struct {
	Construtor IConstrutor
}

func NovoDiretor(c IConstrutor) *Diretor {
	return &Diretor{
		Construtor: c,
	}
}
func (d *Diretor) SetConstrutor(c IConstrutor) {
	d.Construtor = c
}
func (d *Diretor) Construir() Construcao {
	d.Construtor.SetChamine()
	d.Construtor.SetEstiloDeJanelas()
	d.Construtor.SetQuantidadeDeAndares()
	d.Construtor.SetQuantidadeJanelas()
	d.Construtor.SetEstiloDoPiso()
	d.Construtor.SetEstiloDePorta()
	return d.Construtor.GetConstrucao()
}
