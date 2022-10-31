package models

type ConstrutorNormal struct {
	Construcao
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
	c.QuantidadeDeAndares = 2
}
func (c *ConstrutorNormal) SetQuantidadeJanelas() {
	c.QuantidadeDeJanelas = 8
}
func (c *ConstrutorNormal) SetChamine() {
	c.Chamine = "SIM"
}
func (c *ConstrutorNormal) GetConstrucao() Construcao {
	return Construcao{
		TipoDePorta:         c.TipoDePorta,
		TipoDeJanela:        c.TipoDeJanela,
		TipoPiso:            c.TipoPiso,
		Chamine:             c.Chamine,
		QuantidadeDeJanelas: c.QuantidadeDeJanelas,
		QuantidadeDeAndares: c.QuantidadeDeAndares,
	}
}
