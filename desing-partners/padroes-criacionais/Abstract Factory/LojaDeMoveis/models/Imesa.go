package models

type IMesa interface {
	SetDescricao(descricao string)
	SetStyle(style string)
	GetDescricao() string
	GetStyle() string
}

type Mesa struct {
	Descricao string
	Style     string
}

func (c *Mesa) SetDescricao(descricao string) {
	c.Descricao = descricao
}
func (c *Mesa) SetStyle(style string) {
	c.Style = style
}
func (c *Mesa) GetDescricao() string {
	return c.Descricao
}
func (c *Mesa) GetStyle() string {
	return c.Style
}
