package clientes

type Titular struct {
	Nome      string
	Cpf       string
	Profissao string
	Senha     int
}

func (c *Titular) CadSenha(pass int) {
	c.Senha = pass
}
func (c *Titular) VerificarSenha() int {
	return c.Senha
}
func (c *Titular) RegistrarConta() (string, int) {
	return c.Nome, c.Senha
}
