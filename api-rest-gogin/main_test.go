package main

import (
	"api-rest-gogin/controllers"
	"api-rest-gogin/database"
	"api-rest-gogin/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int
var CPF string

func CriaAlunoMock() {
	aluno := models.Aluno{
		Nome: "Aluno teste",
		CPF:  "0123456789X",
		RG:   "12345678X",
	}
	database.DB.Create(&aluno)
	fmt.Println("Criado aluno ID:", aluno.ID)
	ID = int(aluno.ID)
}
func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func SetupDasRotas() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	rotas := gin.Default()
	return rotas
}

func TestStatusCodeDaSaudacaoComParametro(t *testing.T) {
	log.Println("\nTestando saudação:")
	r := SetupDasRotas()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/gui", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	/* if resposta.Code != http.StatusOK {
		t.Fatalf("Status code error: Valor recebido foi %d e o esperado era %d", resposta.Code, http.StatusOK)
	} */
	err := "O código de resposta foi " + strconv.Itoa(resposta.Code) + ", mas deveria ser " + strconv.Itoa(http.StatusOK) + "."
	assert.Equal(t, http.StatusOK, resposta.Code, err)
	mockResposta := `{"Api diz:":"Eai gui, tudo certo!?"}`
	respostaBody, _ := ioutil.ReadAll(resposta.Body)
	assert.Equal(t, mockResposta, string(respostaBody), "A mensagem de retorno não corresponde a configuração.")

}
func TestListandoTodosOsAlunosHandler(t *testing.T) {
	log.Println("\nTestando Listagem de todos os alunos:")
	database.ConectaComDB()
	CriaAlunoMock()
	//defer DeletaAlunoMock()
	r := SetupDasRotas()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
	//fmt.Println(resposta.Body)
}

func TestBuscaAlunoPorCPF(t *testing.T) {
	log.Println("\nTestando a busca de alunos por CPF:")
	database.ConectaComDB()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotas()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoCPF)
	req, _ := http.NewRequest("GET", "/alunos/cpf/0123456789X", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	err := "O código de resposta foi " + strconv.Itoa(resposta.Code) + ", mas deveria ser " + strconv.Itoa(http.StatusOK) + "."
	assert.Equal(t, http.StatusOK, resposta.Code, err)

}
func TestAlunoPorIdHandler(t *testing.T) {
	database.ConectaComDB()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotas()
	r.GET("/alunos/:id", controllers.BuscaAlunoID)
	pathBuscaID := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathBuscaID, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoMock models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)
	/* fmt.Println(alunoMock.Nome)
	fmt.Println(pathBuscaID)
	fmt.Println(resposta.Body) */
	assert.Equal(t, "Aluno teste", alunoMock.Nome, "O nome não corresponde com a posição do DB")
	assert.Equal(t, "0123456789X", alunoMock.CPF, "O CPF não corresponde com a posição do DB")
	assert.Equal(t, "12345678X", alunoMock.RG, "O RG não corresponde com a posição do DB")
	err := "O código de resposta foi " + strconv.Itoa(resposta.Code) + ", mas deveria ser " + strconv.Itoa(http.StatusOK) + "."
	assert.Equal(t, http.StatusOK, resposta.Code, err)

}
func TestDeletaAlunoHandler(t *testing.T) {
	database.ConectaComDB()
	CriaAlunoMock()
	r := SetupDasRotas()
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	pathBuscaID := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathBuscaID, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	err := "O código de resposta foi " + strconv.Itoa(resposta.Code) + ", mas deveria ser " + strconv.Itoa(http.StatusOK) + "."
	assert.Equal(t, http.StatusOK, resposta.Code, err)

}
func TestEditaAlunoHandler(t *testing.T) {
	database.ConectaComDB()
	CriaAlunoMock()
	//defer DeletaAlunoMock()
	r := SetupDasRotas()
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	aluno := models.Aluno{
		Nome: "Aluno",
		CPF:  "12345678901",
		RG:   "123456789",
	}
	pathBuscaID := "/alunos/" + strconv.Itoa(ID)
	alunoJson, _ := json.Marshal(aluno)
	req, _ := http.NewRequest("PATCH", pathBuscaID, bytes.NewBuffer(alunoJson))
	resposta := httptest.NewRecorder()
	fmt.Println("Enviando + bytes:", bytes.NewBuffer(alunoJson))
	r.ServeHTTP(resposta, req)
	var alunoMockEditado models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMockEditado)
	fmt.Println("resp", resposta.Body)

}
