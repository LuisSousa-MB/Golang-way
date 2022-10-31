package models

type Gerente struct {
	Preparador IPreparador
}

func NovoGerente(p IPreparador) *Gerente {
	return &Gerente{
		Preparador: p,
	}
}
func (g *Gerente) SetPreparador(p IPreparador) {
	g.Preparador = p
}
func (g *Gerente) Preparar() Combo {
	g.Preparador.SetNomeCombo()
	g.Preparador.SetHamburguer()
	g.Preparador.SetAcompanhamento()
	g.Preparador.SetBebida()
	g.Preparador.SetSobremesa()
	g.Preparador.SetBrinde()
	g.Preparador.SetDescricao()
	return g.Preparador.GetCombo()
}
