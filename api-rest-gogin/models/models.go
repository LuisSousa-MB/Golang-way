package models

import (
	v "gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	Nome string `json:"nome" validate:"nonzero"`
	CPF  string `json:"cpf" validate:"len=11 ,regexp=^[0-9X-X]*$"`
	RG   string `json:"rg" validate:"len=9 ,regexp=^[0-9X-X]*$"`
}
type AlunoPostReq struct {
	Nome string `json:"nome" validate:"nonzero"`
	CPF  string `json:"cpf" validate:"len=11 ,regexp=^[0-9X-X]*$"`
	RG   string `json:"rg" validate:"len=9 ,regexp=^[0-9X-X]*$"`
}

func ValidaDadosAluno(aluno *Aluno) error {
	if err := v.Validate(aluno); err != nil {
		return err
	}
	return nil
}
