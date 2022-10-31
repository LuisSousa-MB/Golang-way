package models

type PreparadorIceCloudX struct {
	Combo
}

//Hamburguer Z, Suco natural, Batatas-médias, Brinde")
func novoIceCloudX() *PreparadorIceCloudX {
	return &PreparadorIceCloudX{}
}
func (p *PreparadorIceCloudX) SetNomeCombo() {
	p.Nome = "IceCloudX"
}
func (p *PreparadorIceCloudX) SetHamburguer() {
	p.Hamburguer = ""
}
func (p *PreparadorIceCloudX) SetAcompanhamento() {
	p.Acompanhamento = ""
}
func (p *PreparadorIceCloudX) SetBebida() {
	p.Bebida = ""
}
func (p *PreparadorIceCloudX) SetSobremesa() {
	p.Sobremesa = "IceCloudX"
}
func (p *PreparadorIceCloudX) SetBrinde() {
	p.Brinde = ""
}
func (p *PreparadorIceCloudX) SetDescricao() {
	p.SetHamburguer()
	p.SetAcompanhamento()
	p.SetBebida()
	p.SetSobremesa()
	p.Descricao = "Deliciosa sobremesa " + p.Sobremesa + "." + "Sorvete com flocos, Chocolate meio amargo e Frutas vermelhas."
}
func (p *PreparadorIceCloudX) GetCombo() Combo {
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
