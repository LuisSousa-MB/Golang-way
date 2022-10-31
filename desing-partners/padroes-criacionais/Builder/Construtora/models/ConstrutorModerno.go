package models

type ConstrutorModerno struct {
	Construcao
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
	c.QuantidadeDeAndares = 4
}
func (c *ConstrutorModerno) SetQuantidadeJanelas() {
	c.QuantidadeDeJanelas = 12
}
func (c *ConstrutorModerno) SetChamine() {
	c.Chamine = "NÃO"
}
func (c *ConstrutorModerno) GetConstrucao() Construcao {
	return Construcao{
		TipoDePorta:         c.TipoDePorta,
		TipoDeJanela:        c.TipoDeJanela,
		TipoPiso:            c.TipoPiso,
		Chamine:             c.Chamine,
		QuantidadeDeJanelas: c.QuantidadeDeJanelas,
		QuantidadeDeAndares: c.QuantidadeDeAndares,
	}
}
