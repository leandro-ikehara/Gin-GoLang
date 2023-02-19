package main

import (
	"fmt"

	"github.com/leandro-ikehara/Gin-GoLang/database"
	"github.com/leandro-ikehara/Gin-GoLang/routes"
)

func main() {
	database.ConectaBancoDeDados()
	// model.Alunos = []model.Aluno{
	// 	{Nome: "Leandro Ikehara", CPF: "333555999-11", Matricula: "20231020"},
	// 	{Nome: "Ana Paula", CPF: "345126985-77", Matricula: "20231122"},
	// 	{Nome: "Mirella Bastos", CPF: "321654987-88", Matricula: "20230302"},
	// }
	routes.HandleRequests()
	fmt.Println("Iniciando servidor GIN com Go")
}
