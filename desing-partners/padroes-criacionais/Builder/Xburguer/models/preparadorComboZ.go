package models

type PreparadorComboZ struct {
	Combo
}

//Hamburguer Z, Suco natural, Batatas-médias, Brinde")
func novoComboZ() *PreparadorComboZ {
	return &PreparadorComboZ{}
}
func (p *PreparadorComboZ) SetNomeCombo() {
	p.Nome = "Combo Z"
}
func (p *PreparadorComboZ) SetHamburguer() {
	p.Hamburguer = "Hamburguer Z"
}
func (p *PreparadorComboZ) SetAcompanhamento() {
	p.Acompanhamento = "Batatas-médias"
}
func (p *PreparadorComboZ) SetBebida() {
	p.Bebida = "Suco natural"
}
func (p *PreparadorComboZ) SetSobremesa() {
	p.Sobremesa = ""
}
func (p *PreparadorComboZ) SetBrinde() {
	p.Brinde = "Xsummer bubbles"
}
func (p *PreparadorComboZ) SetDescricao() {
	p.SetHamburguer()
	p.SetAcompanhamento()
	p.SetBebida()
	p.SetSobremesa()
	p.Descricao = p.Hamburguer + ", " + p.Acompanhamento + ", " + p.Bebida + " e o super brinde " + p.Brinde + "."
}
func (p *PreparadorComboZ) GetCombo() Combo {
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
