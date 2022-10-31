package prototype

type Inode interface {
	Exibir(string)
	Clonar() Inode
}
