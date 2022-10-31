package models

type PreparadorComboS struct {
	Combo
}

//(Hamburguer S, Suco natural, Salada X, Sobremesa light")
func novoComboS() *PreparadorComboS {
	return &PreparadorComboS{}
}
func (p *PreparadorComboS) SetNomeCombo() {
	p.Nome = "Combo S"
}
func (p *PreparadorComboS) SetHamburguer() {
	p.Hamburguer = "Hamburguer S"
}
func (p *PreparadorComboS) SetAcompanhamento() {
	p.Acompanhamento = "Salada X"
}
func (p *PreparadorComboS) SetBebida() {
	p.Bebida = "Suco natural"
}
func (p *PreparadorComboS) SetSobremesa() {
	p.Sobremesa = "Sobremesa light"
}
func (p *PreparadorComboS) SetBrinde() {
	p.Brinde = ""
}
func (p *PreparadorComboS) SetDescricao() {
	p.SetHamburguer()
	p.SetAcompanhamento()
	p.SetBebida()
	p.SetSobremesa()
	p.Descricao = p.Hamburguer + ", " + p.Acompanhamento + ", " + p.Bebida + ", " + p.Sobremesa + "."
}
func (p *PreparadorComboS) GetCombo() Combo {
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
