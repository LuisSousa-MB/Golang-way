package models

type IPoltrona interface {
	SetDescricao(descricao string)
	SetStyle(style string)
	GetDescricao() string
	GetStyle() string
}

type Poltrona struct {
	Descricao string
	Style     string
}

func (c *Poltrona) SetDescricao(descricao string) {
	c.Descricao = descricao
}
func (c *Poltrona) SetStyle(style string) {
	c.Style = style
}
func (c *Poltrona) GetDescricao() string {
	return c.Descricao
}
func (c *Poltrona) GetStyle() string {
	return c.Style
}
