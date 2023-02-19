package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/leandro-ikehara/Gin-GoLang/controller"
	"github.com/leandro-ikehara/Gin-GoLang/database"
	"github.com/leandro-ikehara/Gin-GoLang/model"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock() {
	aluno := model.Aluno{Nome: "Aluno Teste", CPF: "789465132377", Matricula: "789456123"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno model.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestVerificaStatusCode(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controller.Saudacao)
	req, _ := http.NewRequest("GET", "/leandro", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais")
	mockDaResposta := `{"API diz":"Olá, leandro, traquilo?"}`
	respostaBody, _ := ioutil.ReadAll(resposta.Body)
	assert.Equal(t, mockDaResposta, string(respostaBody))
}

func TestListandoAlunosHandler(t *testing.T) {
	database.ConectaBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controller.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais")
	// fmt.Println((resposta.Body))
}

func TestBuscaAlunoPorCpfHandler(t *testing.T) {
	database.ConectaBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/cpf/:cpf", controller.BuscaAlunoPorCpf)
	req, _ := http.NewRequest("GET", "/alunos/cpf/789465132377", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}
