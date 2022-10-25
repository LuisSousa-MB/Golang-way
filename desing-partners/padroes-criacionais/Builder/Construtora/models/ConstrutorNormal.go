package models

type ConstrutorNormal struct {
	TipoDeJanela, TipoDePorta, Chamine, TipoPiso string
	Andares, Janelas                             int
}

func novoConstrutorNormal() *ConstrutorNormal {
	return &ConstrutorNormal{}
}
func (c *ConstrutorNormal) SetEstiloDeJanelas() {
	c.TipoDeJanela = "Vidraça normal" // implementar
}
func (c *ConstrutorNormal) SetEstiloDePorta() {
	c.TipoDePorta = "Madeira" // implementar
}
func (c *ConstrutorNormal) SetEstiloDoPiso() {
	c.TipoPiso = "Porcelanato" // implementar
}
func (c *ConstrutorNormal) SetQuantidadeDeAndares() {
	c.Andares = 2
}
func (c *ConstrutorNormal) SetQuantidadeJanelas() {
	c.Janelas = 8
}
func (c *ConstrutorNormal) SetChamine(adicionar bool) {
	if adicionar == true {
		c.Chamine = "SIM"
	} else {
		c.Chamine = "NÃO"
	}
}
func (c *ConstrutorNormal) GetConstrucao() Construcao {
	return Construcao{
		TipoDePorta:         c.TipoDePorta,
		TipoDeJanela:        c.TipoDeJanela,
		TipoPiso:            c.TipoPiso,
		Chamine:             c.Chamine,
		QuantidadeDeJanelas: c.Janelas,
		QuantidadeDeAndares: c.Andares,
	}
}
