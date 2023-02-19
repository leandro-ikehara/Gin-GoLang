package model

import "gorm.io/gorm"

type Aluno struct {
	gorm.Model
	Nome      string `json:"nome"`
	CPF       string `json:"cpf"`
	Matricula string `json:"matricula"`
}

// var Alunos []Aluno
