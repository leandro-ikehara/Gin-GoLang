package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/leandro-ikehara/Gin-GoLang/controller"
	"github.com/stretchr/testify/assert"
)

func SetupDasRotasDeTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func TestVerificaStatusCode(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controller.Saudacao)
	req, _ := http.NewRequest("GET", "/leandro", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	// if resposta.Code != http.StatusOK {
	// 	t.Fatalf("Status error: valor recebido foi %d e o esperado era %d", resposta.Code, http.StatusOK)
	// }
	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais")
	mockDaResposta := `{"API diz":"Olá, leandro, traquilo?"}`
	respostaBody, _ := ioutil.ReadAll(resposta.Body)
	assert.Equal(t, mockDaResposta, string(respostaBody))
	// fmt.Println(string(respostaBody))
	// fmt.Println(mockDaResposta)
}
