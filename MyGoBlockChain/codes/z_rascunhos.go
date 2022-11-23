package codes

import (
	"fmt"
	"time"
)

func VerificarSenhaSecreta(pass string) bool {
	if pass == "wEBHqEHkPoMngxeGBWMEaFcAfXEwfhKR" {
		return true
	} else {
		return false
	}
}

func GoCrypter() {
	//var pass string
	NewChain.GerateGenesisBlock()
	fmt.Println("Informe a senha:")
	//fmt.Scan(&pass)
	time.Sleep(time.Second)
	if VerificarSenhaSecreta(GenerateHashByte("Estamos indo para a Lua!")) {
		println("Autorizado!")
	} else {
		println("Acesso negado!")
	}
}
