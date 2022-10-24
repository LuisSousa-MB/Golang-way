package models

type ICadeira interface {
	SetDescricao(descricao string)
	SetStyle(style string)
	GetDescricao() string
	GetStyle() string
}

type Cadeira struct {
	Descricao string
	Style     string
}

func (c *Cadeira) SetDescricao(descricao string) {
	c.Descricao = descricao
}
func (c *Cadeira) SetStyle(style string) {
	c.Style = style
}
func (c *Cadeira) GetDescricao() string {
	return c.Descricao
}
func (c *Cadeira) GetStyle() string {
	return c.Style
}
