package models

type PreparadorComboX struct {
	Combo
}

func novoComboX() *PreparadorComboX {
	return &PreparadorComboX{}
}
func (p *PreparadorComboX) SetNomeCombo() {
	p.Nome = "Combo X"
}
func (p *PreparadorComboX) SetHamburguer() {
	p.Hamburguer = "Hamburguer X"
}
func (p *PreparadorComboX) SetAcompanhamento() {
	p.Acompanhamento = "Batatas-médias"
}
func (p *PreparadorComboX) SetBebida() {
	p.Bebida = "Refrigerante"
}
func (p *PreparadorComboX) SetSobremesa() {
	p.Sobremesa = "Sobremesa do dia"
}
func (p *PreparadorComboX) SetBrinde() {
	p.Brinde = "Xsummer pistol"
}
func (p *PreparadorComboX) SetDescricao() {
	p.SetHamburguer()
	p.SetAcompanhamento()
	p.SetBebida()
	p.SetSobremesa()
	p.SetBrinde()
	x := p.Hamburguer + ", " + p.Acompanhamento + ", " + p.Bebida + ", " + p.Sobremesa + " e o super brinde " + p.Brinde + "."
	p.Descricao = x
}
func (p *PreparadorComboX) GetCombo() Combo {
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
