package models

import "gorm.io/gorm"

type Titular struct {
	gorm.Model
	Nome      string `json:"nome" validate:"nonzero"`
	Cpf       string `json:"cpf" validate:"nonzero"`
	Profissao string `json:"profissao" validate:"nonzero"`
	Senha     int    `json:"senha" validate:"nonzero"`
}

type ContaCorrente struct {
	gorm.Model
	Titular                        Titular `json:"location,omitempty" gorm:"foreignKey:LocationID;references:ID"`
	NumAgencia, NumConta, Operacao int
	saldo                          float64
}

type ContaPoupanca struct {
	gorm.Model
	Titular                        int
	NumAgencia, NumConta, Operacao int
	saldo                          float64
}
