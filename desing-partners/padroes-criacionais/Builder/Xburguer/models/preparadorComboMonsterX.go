package models

type PreparadorComboMonsterX struct {
	Combo
}

func novoComboMonsterX() *PreparadorComboMonsterX {
	return &PreparadorComboMonsterX{}
}
func (p *PreparadorComboMonsterX) SetNomeCombo() {
	p.Nome = "Combo Monster X"
}
func (p *PreparadorComboMonsterX) SetHamburguer() {
	p.Hamburguer = "Hamburguer TripleX"
}
func (p *PreparadorComboMonsterX) SetAcompanhamento() {
	p.Acompanhamento = "Batatas-grandes"
}
func (p *PreparadorComboMonsterX) SetBebida() {
	p.Bebida = "Refrigerante-XBig"
}
func (p *PreparadorComboMonsterX) SetSobremesa() {
	p.Sobremesa = "IceCouldx"
}
func (p *PreparadorComboMonsterX) SetBrinde() {
	p.Brinde = "Xsummer pistol e Xsummer bubbles"
}
func (p *PreparadorComboMonsterX) SetDescricao() {
	p.SetHamburguer()
	p.SetAcompanhamento()
	p.SetBebida()
	p.SetSobremesa()
	p.SetBrinde()
	p.Descricao = p.Hamburguer + ", " + p.Acompanhamento + ", " + p.Bebida + ", " + p.Sobremesa + " e os super brindes " + p.Brinde + "."
}
func (p *PreparadorComboMonsterX) GetCombo() Combo {
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
