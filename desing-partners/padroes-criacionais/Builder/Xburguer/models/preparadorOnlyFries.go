package models

type PreparadorOnlyFries struct {
	Combo
}

//Hamburguer Z, Suco natural, Batatas-médias, Brinde")
func novoOnlyFries() *PreparadorOnlyFries {
	return &PreparadorOnlyFries{}
}
func (p *PreparadorOnlyFries) SetNomeCombo() {
	p.Nome = "Only Fries"
}
func (p *PreparadorOnlyFries) SetHamburguer() {
	p.Hamburguer = ""
}
func (p *PreparadorOnlyFries) SetAcompanhamento() {
	p.Acompanhamento = "Batatas-média"
}
func (p *PreparadorOnlyFries) SetBebida() {
	p.Bebida = ""
}
func (p *PreparadorOnlyFries) SetSobremesa() {
	p.Sobremesa = ""
}
func (p *PreparadorOnlyFries) SetBrinde() {
	p.Brinde = ""
}
func (p *PreparadorOnlyFries) SetDescricao() {
	p.SetHamburguer()
	p.SetAcompanhamento()
	p.SetBebida()
	p.SetSobremesa()
	p.Descricao = p.Acompanhamento + "."
}
func (p *PreparadorOnlyFries) GetCombo() Combo {
	return Combo{
		Nome:           p.Nome,
		Hamburguer:     p.Hamburguer,
		Acompanhamento: p.Acompanhamento,
		Bebida:         p.Bebida,
		Sobremesa:      p.Sobremesa,
		Brinde:         p.Brinde,
		Descricao:      p.Descricao,
	}

}
