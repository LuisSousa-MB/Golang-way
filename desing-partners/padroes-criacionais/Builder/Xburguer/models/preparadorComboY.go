package models

type PreparadorComboY struct {
	Combo
}

//((Hamburguer Y, Refrigerante, Batatas-grande, Sobremesa do dia")
func novoComboY() *PreparadorComboY {
	return &PreparadorComboY{}
}
func (p *PreparadorComboY) SetNomeCombo() {
	p.Nome = "Combo Y"
}

func (p *PreparadorComboY) SetHamburguer() {
	p.Hamburguer = "Hamburguer Y"
}
func (p *PreparadorComboY) SetAcompanhamento() {
	p.Acompanhamento = "Batatas-grande"
}
func (p *PreparadorComboY) SetBebida() {
	p.Bebida = "Refrigerante"
}
func (p *PreparadorComboY) SetSobremesa() {
	p.Sobremesa = "Sobremesa do dia"
}
func (p *PreparadorComboY) SetBrinde() {
	p.Brinde = ""
}
func (p *PreparadorComboY) SetDescricao() {
	p.SetHamburguer()
	p.SetAcompanhamento()
	p.SetBebida()
	p.SetSobremesa()
	p.Descricao = p.Hamburguer + ", " + p.Acompanhamento + ", " + p.Bebida + ", " + p.Sobremesa + "."
}
func (p *PreparadorComboY) GetCombo() Combo {
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
