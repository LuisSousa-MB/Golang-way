package models

type ICama interface {
	SetDescricao(descricao string)
	SetStyle(style string)
	GetDescricao() string
	GetStyle() string
}

type Cama struct {
	Descricao string
	Style     string
}

func (c *Cama) SetDescricao(descricao string) {
	c.Descricao = descricao
}
func (c *Cama) SetStyle(style string) {
	c.Style = style
}
func (c *Cama) GetDescricao() string {
	return c.Descricao
}
func (c *Cama) GetStyle() string {
	return c.Style
}
