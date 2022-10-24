package models

type IArmario interface {
	SetDescricao(descricao string)
	SetStyle(style string)
	GetDescricao() string
	GetStyle() string
}

type Armario struct {
	Descricao string
	Style     string
}

func (c *Armario) SetDescricao(descricao string) {
	c.Descricao = descricao
}
func (c *Armario) SetStyle(style string) {
	c.Style = style
}
func (c *Armario) GetDescricao() string {
	return c.Descricao
}
func (c *Armario) GetStyle() string {
	return c.Style
}
