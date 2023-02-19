package model

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	Nome      string `json:"nome" validate:"nonzero"`
	CPF       string `json:"cpf" validate:"len=11, regexp=^[0-9]*$"`
	Matricula string `json:"matricula" validate:"len=9, regexp=^[0-9]*$"`
}

func ValidaDadosDeAluno(aluno *Aluno) error {
	if err := validator.Validate(aluno); err != nil {
		return err
	}
	return nil
}
