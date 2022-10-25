package models

type ConstrutorModerno struct {
	TipoDeJanela, TipoDePorta, Chamine, TipoPiso string
	Andares, Janelas                             int
}

func novoConstrutorModerno() *ConstrutorModerno {
	return &ConstrutorModerno{}
}
func (c *ConstrutorModerno) SetEstiloDeJanelas() {
	c.TipoDeJanela = "Espelhadas" // implementar
}
func (c *ConstrutorModerno) SetEstiloDePorta() {
	c.TipoDePorta = "Aço importado" // implementar
}
func (c *ConstrutorModerno) SetEstiloDoPiso() {
	c.TipoPiso = "Marmore indiano" // implementar
}
func (c *ConstrutorModerno) SetQuantidadeDeAndares() {
	c.Andares = 4
}
func (c *ConstrutorModerno) SetQuantidadeJanelas() {
	c.Janelas = 12
}
func (c *ConstrutorModerno) SetChamine(adicionar bool) {
	if adicionar == true {
		c.Chamine = "SIM"
	} else {
		c.Chamine = "NÃO"
	}
}
func (c *ConstrutorModerno) GetConstrucao() Construcao {
	return Construcao{
		TipoDePorta:         c.TipoDePorta,
		TipoDeJanela:        c.TipoDeJanela,
		TipoPiso:            c.TipoPiso,
		Chamine:             c.Chamine,
		QuantidadeDeJanelas: c.Janelas,
		QuantidadeDeAndares: c.Andares,
	}
}
